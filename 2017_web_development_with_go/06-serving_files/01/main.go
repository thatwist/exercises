// ListenAndServe on port 8080 of localhost
// For the default route "/" Have a func called "foo" which writes to the response "foo ran"
// For the route "/dog/" Have a func called "dog" which parses a template called "dog.gohtml" and writes to the response "
// This is from dog
// " and also shows a picture of a dog when the template is executed.
// Use "http.ServeFile" to serve the file "dog.jpeg"

package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "foo run")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.Execute(w, nil)
}

func dogImg(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "dog.jpeg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/dog.jpeg", dogImg)
	http.ListenAndServe(":8080", nil)
}
