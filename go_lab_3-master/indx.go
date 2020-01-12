package main
import "net/http"
	   "html/template"

var templ = template.Must(template.ParseFiles("indexx.html"))

func index (w http.ResponseWriter, r *http.Request){
	templ.Execute(w,nil)

}

func main (){
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexx)
	http.ListenAndServe(":8080",mux)
}