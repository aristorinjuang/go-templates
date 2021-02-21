package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template
var response bytes.Buffer

func init() {
	tmpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	if err := tmpl.ExecuteTemplate(&response, "index.gohtml", "Nested"); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, response.String())
	})
	http.ListenAndServe(":80", nil)
}
