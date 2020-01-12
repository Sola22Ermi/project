package main

import (
	"database/sql"
	"fmt"
	"github.com"
	"html/template"
	"net/http"
	"strconv"
)
var db *sql.DB
var tpl *template.Template



func init() {
	var err error
	db, err = sql.Open("postgres","postgres://bond:password@localhost/bookList?sslmode=disable")
	if err !=nil{
		panic (err)
	}

	if err =db.Ping(); err != nil{
		panic(err)
	}
	fmt.Println("you connected to your database.")

	tpl = template.Must(template.ParseGlob("templates/*."))//html.file

}

type Book struct{
	Isbn string
	Title string
	Author string
	pirce float32
}
}
func main() {
	http.HandleFunc("/",index)
	http.HandleFunc("/books",booksIndex)
	http.HandleFunc("/books/show", booksShow)
	http.HandleFunc("/books/create/process",booksCreateProcess)

}

func index (w http.ResponseWriter,r *http.Request){
	http.Redirect(w,r,"/books", http.StatusSeeOther)
}

func booksIndex(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
		http.Error(w,http.StatusText(405),http.Status)
	}
}
func booksDeleteProcess(w http.ResponseWriter,r*http.Request){
	if r.Method != "GET"{
		http.Error(w,http.StatusText(405),http.StatusMethodNotAllowed)
		return
	}
	isbn :=r.FormValue("isbn")
	if isbn ==""{
		http.Error(w,http.StatusText(400), http.StatusBadRequest)
		return
	}
	// delete book
	_, err := db.Exec("")
}

	
