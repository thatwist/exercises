// Take the previous program in the previous folder and change it so that:
// a template is parsed and served
// you pass data into the template.

package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("layout.gohtml"))
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	tpl.Execute(w, "Hello World")
}

func handleDog(w http.ResponseWriter, req *http.Request) {
	tpl.Execute(w, "av av av")
}

func handleMe(w http.ResponseWriter, req *http.Request) {
	tpl.Execute(w, "Dejan")
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/me/", handleMe)
	http.ListenAndServe(":8080", nil)
}
