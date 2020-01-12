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

	// alternatively
	// http.HandleFunc("/index", indexHandler)
	// http.HandleFunc("/contact", contactHandler)
	// http.ListenAndServe(":8080", nil)
	// //it defaults to DefaultServeMux.
}

// ------ 01 start server

// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// var templ = template.Must(template.ParseFiles("index.html"))

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	templ.Execute(w, nil)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", indexHandler)
// 	http.ListenAndServe(":8080", mux)
// }
// // --------02 using template

// import (
// 	"html/template"
// 	"net/http"
// )

// // Course struct
// type Course struct {
// 	Title, Code, Description string
// 	ECTS                     int
// }

// var templ = template.Must(template.ParseFiles("class.html"))

// func classHandler(w http.ResponseWriter, r *http.Request) {
// 	crs := Course{"Web programing", "ITSE-3182", "Lorem Ipsum", 7}
// 	templ.Execute(w, crs)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/class", classHandler)
// 	http.ListenAndServe(":8080", mux)
// }

// // --------03 passing data to template

// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// var templ = template.Must(template.ParseFiles("indexstyle.html"))

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	templ.Execute(w, nil)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	fs := http.FileServer(http.Dir("assets"))
// 	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
// 	mux.HandleFunc("/", indexHandler)
// 	http.ListenAndServe(":8080", mux)
// }

// // ----- 04 serving static files

// package main

// import (
// 	"html/template"
// 	"net/http"
// )

// var templ = template.Must(template.ParseFiles("bootstraped.html"))

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	templ.Execute(w, nil)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	fs := http.FileServer(http.Dir("assets"))
// 	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))
// 	mux.HandleFunc("/", indexHandler)
// 	http.ListenAndServe(":8080", mux)
// }

// ----- 05 using bootstrap
