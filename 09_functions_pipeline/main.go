package main

import (
	"bytes"
	"io"
	"log"
	"math"
	"net/http"
	"text/template"
)

func one(number int) int {
	return int(math.Pow(float64((1+math.Sqrt(5))/2), float64(number)))
}

func two(number int) int {
	return int(math.Round(float64(number) / math.Sqrt(5)))
}

var tmpl *template.Template
var response bytes.Buffer

var functions = template.FuncMap{
	"one": one,
	"two": two,
}

func init() {
	tmpl = template.Must(template.New("tmpl.gohtml").Funcs(functions).ParseFiles("tmpl.gohtml"))
}

func main() {
	if err := tmpl.Execute(&response, 7); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, response.String())
	})
	http.ListenAndServe(":80", nil)
}
