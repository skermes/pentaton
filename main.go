package main

import (
  "os"
  "io"
  "fmt"
  "net/http"
  "html/template"

  "github.com/zenazn/goji"
  "github.com/zenazn/goji/web"
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

func main() {
  setup()
  goji.Get("/hello/:name", hello)
  goji.Serve()
}
