package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"sort"

	"github.com/skermes/pentaton/Godeps/_workspace/src/github.com/zenazn/goji"
	"github.com/skermes/pentaton/Godeps/_workspace/src/github.com/zenazn/goji/web"

	_ "github.com/skermes/pentaton/Godeps/_workspace/src/github.com/lib/pq"
)

var templates *template.Template
var err error
var db *sql.DB

func setup() {
	templates, err = template.ParseGlob("templates/*")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error loading templates: %s", err.Error())
	}

	db, err = sql.Open("postgres", "postgres://skermes:skermes@localhost:5432/pentaton?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error opening database connection: %s", err.Error())
	}
}

func render(w io.Writer, tmpl string, data interface{}) {
	err = templates.ExecuteTemplate(w, tmpl, data)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error executing template '%s': %s", tmpl, err.Error())
	}
}

type Link struct {
	Url string
	Name string
	Color string
	Position int
	Category string
}

type ByPosition []Link

// Implement sort.Interface
func (a ByPosition) Len() int           { return len(a) }
func (a ByPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPosition) Less(i, j int) bool { return a[i].Position < a[j].Position }

func links(c web.C, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`
		select url, links.name, color, position, categories.name from
		links
		join categories
		on links.category = categories.id
	`)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error reading from database: %s", err.Error())
		return
	}

	links := make([]Link, 0)

	for rows.Next() {
		var url, name, color, category string
		var position int
		if err := rows.Scan(&url, &name, &color, &position, &category); err != nil {
			fmt.Fprintf(os.Stdout, "Error reading row: %s", err.Error())
		}

		links = append(links, Link{url, name, color, position, category})
	}

	sort.Sort(ByPosition(links))
	render(w, "links", links)
}

func main() {
	setup()
	goji.Get("/", links)
	goji.Serve()
}
