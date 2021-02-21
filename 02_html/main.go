package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
)

var tmpl *template.Template
var response bytes.Buffer

func init() {
	tmpl = template.Must(template.ParseFiles("hello.gohtml"))
}

func main() {
	if err := tmpl.Execute(&response, `<script>alert("Hello World!")</script>`); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, response.String())
	})
	http.ListenAndServe(":80", nil)
}
