// Building upon the code from the previous problem:
// Add code to respond to the following METHODS & ROUTES: GET / GET /apply POST /apply

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			continue
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	ln := scanner.Text()
	fmt.Println(ln)
	fields := strings.Fields(ln)
	if len(fields) < 2 {
		respondBadRequest(conn, "Unknown protocol")
		return
	}
	method := fields[0]
	uri := fields[1]
	handle(conn, method, uri)
}

type request struct {
	method string
	uri    string
}

func handle(conn net.Conn, method, uri string) {
	a := request{method, uri}
	switch a {
	case request{"GET", "/"}:
		respondOk(conn, "GETTING INDEX")
	case request{"GET", "/apply"}:
		respondOk(conn, "GETTING APPLY")
	case request{"POST", "/apply"}:
		respondOk(conn, "POSTING APPLY")
	default:
		respondNotFound(conn, "PAGE NOT FOUND")
	}
}

func respondOk(conn net.Conn, body string) {
	respond(conn, body, "200 OK")
}

func respondNotFound(conn net.Conn, body string) {
	respond(conn, body, "404 NOT FOUND")
}

func respondBadRequest(conn net.Conn, body string) {
	respond(conn, body, "400 BAD REQUEST")
}

func respond(conn net.Conn, body, status string) {
	fmt.Fprintf(conn, "HTTP/1.1 %s\r\n", status)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
