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
	http.Handle("/", http.HandlerFunc(handleIndex))
	http.Handle("/dog/", http.HandlerFunc(handleDog))
	http.Handle("/me/", http.HandlerFunc(handleMe))
	http.ListenAndServe(":8080", nil)
}
