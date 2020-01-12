package main

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("homepage.html"))

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}

func main() {
	mux := http.NewServeMux()
	http.HandleFunc("/", index)
	http.ListenAndServe("localhost:8080", mux)

}
