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
	tmpl = template.Must(template.ParseFiles("tmpl.gohtml"))
}

func main() {
	data := struct {
		Numbers []int
		Arg1    int
		Arg2    int
		Arg3    int
	}{
		[]int{1, 2, 3},
		5,
		7,
		5,
	}
	if err := tmpl.Execute(&response, data); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, response.String())
	})
	http.ListenAndServe(":80", nil)
}
