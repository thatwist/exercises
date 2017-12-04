// create a mux so that your server responds to different URIs and METHODS

package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	if len(fields) < 2 {
		respondBadRequest(conn, "Unknown protocol")
		return
	}

	switch fields[0] {
	case "GET":
		handleGet(conn, fields[1])
	case "POST":
		respondNotFound(conn)
	default:
		respondBadRequest(conn, "Unknown method")
	}
}

func handleGet(w io.Writer, path string) {
	switch path {
	case "/":
		respondOk(w, "index page")
	case "/login":
		respondOk(w, "login page")
	default:
		respondNotFound(w)
	}
}

func respond(w io.Writer, status string, body string) {
	fmt.Fprintf(w, "HTTP/1.1 %s\r\n", status)
	fmt.Fprintf(w, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(w, "\r\n")
	fmt.Fprint(w, body)
}

func respondOk(w io.Writer, body string) {
	respond(w, "200 OK", body)
}

func respondBadRequest(w io.Writer, body string) {
	respond(w, "400 BAD REQUEST", body)
}

func respondNotFound(w io.Writer) {
	respond(w, "404 NOT FOUND", "Page not found")
}
