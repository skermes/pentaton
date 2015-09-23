package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/skermes/pentaton/Godeps/_workspace/src/github.com/zenazn/goji"
	"github.com/skermes/pentaton/Godeps/_workspace/src/github.com/zenazn/goji/web"

	_ "github.com/skermes/pentaton/Godeps/_workspace/src/github.com/lib/pq"
)

var templates *template.Template
var err error

func setup() {
	templates, err = template.ParseGlob("templates/*")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error loading templates: %s", err.Error())
	}
}

func render(w io.Writer, tmpl string, data interface{}) {
	err = templates.ExecuteTemplate(w, tmpl, data)

	if err != nil {
		fmt.Fprintf(os.Stdout, "Error executing template '%s': %s", tmpl, err.Error())
	}
}

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	render(w, "hello", map[string]string{"Name": c.URLParams["name"]})
}

func row(c web.C, w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "postgres://skermes:skermes@localhost:5432/pentaton?sslmode=disable")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error connecting to database: %s", err.Error())
		return
	}

	rows, err := db.Query("select name, url, color from links limit 1")
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error reading from database: %s", err.Error())
		return
	}

	rows.Next()

	var name, url, color string
	err = rows.Scan(&name, &url, &color)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error scanning rows: %s", err.Error())
		return
	}

	render(w, "db", map[string]string{"Name": name, "Url": url, "Color": color})
}

func main() {
	setup()
	goji.Get("/hello/:name", hello)
	goji.Get("/row", row)
	goji.Serve()
}
