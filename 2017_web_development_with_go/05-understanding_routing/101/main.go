// ListenAndServe on port ":8080" using the default ServeMux.
// Use HandleFunc to add the following routes to the default ServeMux:
// "/" "/dog/" "/me/
// Add a func for each of the routes.
// Have the "/me/" route print out your name.

package main

import (
	"fmt"
	"net/http"
)

func handleIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func handleDog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "av av av")
}

func handleMe(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Dejan")
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/dog/", handleDog)
	http.HandleFunc("/me/", handleMe)
	http.ListenAndServe(":8080", nil)
}
