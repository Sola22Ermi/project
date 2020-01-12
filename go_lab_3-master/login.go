package main

import "net/http"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hello world</h1>"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hii</h1>"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/contact", contactHandler)
	http.ListenAndServe(":8080", mux)
