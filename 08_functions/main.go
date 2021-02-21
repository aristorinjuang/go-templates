package main

import (
	"bytes"
	"io"
	"log"
	"math"
	"net/http"
	"strings"
	"text/template"
)

func fibonacci(number int) int {
	return int(math.Round(math.Pow(float64((1+math.Sqrt(5))/2), float64(number)) / math.Sqrt(5)))
}

var tmpl *template.Template

var functions = template.FuncMap{
	"title":     strings.Title,
	"fibonacci": fibonacci,
}

func init() {
	tmpl = template.Must(template.New("tmpl.gohtml").Funcs(functions).ParseFiles("tmpl.gohtml"))
}

func main() {
	var response bytes.Buffer
	data := struct {
		Names   []string
		Numbers []int
	}{
		[]string{"john doe", "john cena", "harry potter"},
		[]int{1, 3, 5, 7, 9},
	}
	if err := tmpl.Execute(&response, data); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, response.String())
	})
	http.ListenAndServe(":80", nil)
}
