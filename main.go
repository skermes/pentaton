package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/skermes/pentaton/Godeps/_workspace/src/github.com/zenazn/goji"
	"github.com/skermes/pentaton/Godeps/_workspace/src/github.com/zenazn/goji/web"

	_ "github.com/skermes/pentaton/Godeps/_workspace/src/github.com/lib/pq"
)

const numColumns = 3

var templates *template.Template
var err error
var db *sql.DB

func setup() {
	funcs := template.FuncMap{
		"mod": func(x, y int) int { return x % y },
		"rgbstr": func(hex string) string {
			rgb, err := strconv.ParseInt(hex, 16, 32)
			if err != nil {
				fmt.Printf("Error parsing color string '%s': %s", hex, err.Error())
				return ""
			}

			r := (rgb & 0xFF0000) >> 16
			g := (rgb & 0xFF00) >> 8
			b := rgb & 0xFF

			return fmt.Sprintf("%d, %d, %d", r, g, b)
		},
	}

	templates, err = template.New("").Funcs(funcs).ParseGlob("templates/*")
	if err != nil {
		fmt.Printf("Error loading templates: %s", err.Error())
	}

	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
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

// This is kind of a gross hack, but it lets me add things to the list of
// links that aren't links, like a marker for where the new link form should
// go.
type LinkWidget struct {
	Url string
	Name string
	Color string
	Position int

	Widget string
	Category string
}

type ByPosition []LinkWidget

// Implement sort.Interface
func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

func getPartitionedLinks(category string) [][]LinkWidget {
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

	links := make([]LinkWidget, 0)

	for rows.Next() {
		var url, name, color string
		var position int
		if err := rows.Scan(&url, &name, &color, &position); err != nil {
			fmt.Printf("Error reading row: %s", err.Error())
			return nil
		}

		links = append(links, LinkWidget{Url: url, Name: name, Color: color, Position: position})
	}

	sort.Sort(ByPosition(links))
	links = append(links, LinkWidget{Widget: "new-link-form", Category: category})

	partitioned := make([][]LinkWidget, 0)
	for i := 0; i < len(links); i++ {
		row := i / numColumns

		if len(partitioned) < row + 1 {
			partitioned = append(partitioned, make([]LinkWidget, 0))
		}

		partitioned[row] = append(partitioned[row], links[i])
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

func newLink(c web.C, w http.ResponseWriter, r *http.Request) {
	category := c.URLParams["category"]
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

	color := r.PostFormValue("color")
	url := r.PostFormValue("url")
	name := r.PostFormValue("name")

	_, err := db.Exec(`
		insert into links
		(color, url, name, position, category)
		values
		($1, $2, $3,
		 (select max(position) from links) + 1,
		 (select id from categories where name = $4))
	`, color, url, name, category)

	if err != nil {
		fmt.Println("Error saving new link: %s", err.Error())
	}

	linksForCategory(c, w, r, category)
}

func main() {
	setup()
	goji.Get("/", readingLinks)
	goji.Get("/:category", links)
	goji.Post("/:category", newLink)
	goji.Serve()
}
