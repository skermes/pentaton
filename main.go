package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"sort"

	"github.com/skermes/pentaton/Godeps/_workspace/src/github.com/zenazn/goji"
	"github.com/skermes/pentaton/Godeps/_workspace/src/github.com/zenazn/goji/web"

	_ "github.com/skermes/pentaton/Godeps/_workspace/src/github.com/lib/pq"
)

const numColumns = 3

var templates *template.Template
var err error
var db *sql.DB

func setup() {
	templates, err = template.ParseGlob("templates/*")
	if err != nil {
		fmt.Printf("Error loading templates: %s", err.Error())
	}

	db, err = sql.Open("postgres", "postgres://skermes:skermes@localhost:5432/pentaton?sslmode=disable")
	if err != nil {
		fmt.Printf("Error opening database connection: %s", err.Error())
	}
}

func render(w io.Writer, tmpl string, data interface{}) {
	err = templates.ExecuteTemplate(w, tmpl, data)

	if err != nil {
		fmt.Printf("Error executing template '%s': %s", tmpl, err.Error())
	}
}

type Link struct {
	Url string
	Name string
	Color string
	Position int
}

type ByPosition []Link

// Implement sort.Interface
func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

func getPartitionedLinks(category string) [][]Link {
	rows, err := db.Query(`
		select url, links.name, color, position from
		links
		join categories
		on links.category = categories.id
		where categories.name = $1
	`, category)
	if err != nil {
		fmt.Printf("Error reading from database: %s", err.Error())
		return nil
	}

	links := make([]Link, 0)

	for rows.Next() {
		var url, name, color string
		var position int
		if err := rows.Scan(&url, &name, &color, &position); err != nil {
			fmt.Printf("Error reading row: %s", err.Error())
			return nil
		}

		links = append(links, Link{url, name, color, position})
	}

	sort.Sort(ByPosition(links))

	partitioned := make([][]Link, 0)
	for i := 0; i < len(links); i++ {
		row := i / numColumns
		col := i % numColumns

		if len(partitioned) < row + 1 {
			partitioned = append(partitioned, make([]Link, numColumns))
		}

		partitioned[row][col] = links[i];
	}

	return partitioned
}

func getCategories() []string {
	rows, err := db.Query(`select name from categories`)
	if err != nil {
		fmt.Printf("Error reading from database: %s", err.Error())
		return nil
	}

	categories := make([]string, 0)

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			fmt.Printf("Error reading row: %s", err.Error())
			return nil
		}

		categories = append(categories, category)
	}

	return categories
}

func linksForCategory(c web.C, w http.ResponseWriter, r *http.Request, category string) {
	categories := getCategories()

	contains := false
	for _, c := range(categories) {
		if c == category {
			contains = true
		}
	}

	if !contains {
		w.WriteHeader(http.StatusNotFound)
		render(w, "404", map[string]interface{}{
			"Category": category,
			"Categories": categories,
		})
		return
	}

	render(w, "links", map[string]interface{}{
		"Links": getPartitionedLinks(category),
		"Category": category,
		"Categories": categories,
	})
}

func links(c web.C, w http.ResponseWriter, r *http.Request) {
	linksForCategory(c, w, r, c.URLParams["category"])
}

func readingLinks(c web.C, w http.ResponseWriter, r *http.Request) {
	linksForCategory(c, w, r, "reading")
}

func main() {
	setup()
	goji.Get("/", readingLinks)
	goji.Get("/:category", links)
	goji.Serve()
}
