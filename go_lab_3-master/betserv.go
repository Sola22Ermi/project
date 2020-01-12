package main

import (
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1>man what's up</h1>"))

}

func main() {
	

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)

}
