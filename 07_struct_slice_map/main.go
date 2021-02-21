package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	var response bytes.Buffer
	data := struct {
		Names []string
		Jobs  map[string]string
	}{
		[]string{"John Doe", "John Cena", "Harry Potter"},
		map[string]string{
			"Dummy":    "John Doe",
			"Wrestler": "John Cena",
			"Magician": "Harry Potter",
		},
	}
	if err := tmpl.Execute(&response, data); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, response.String())
	})
	http.ListenAndServe(":80", nil)
}
