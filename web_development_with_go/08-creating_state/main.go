// Using cookies, track how many times a user has been to your website domain.

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("counter")
	var counts int

	if err == nil {
		prevVisits, err := strconv.Atoi(cookie.Value)
		if err == nil {
			counts = prevVisits + 1
		}
	} else {
		cookie = &http.Cookie{Name: "counter"}
		counts = 0
	}

	cookie.Value = strconv.Itoa(counts)
	http.SetCookie(w, cookie)
	fmt.Fprintf(w, "Hi there! You were here before exactly: %d times", counts)
}
