// create a HTTP server without using http package that returns URL of the GET request

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
	defer closeConn(conn)
	fmt.Println("New connection...")
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	ln := scanner.Text()
	f := strings.Fields(ln)
	if len(f) > 2 {
		respond(conn, "200 OK", f[1])
	} else {
		respond(conn, "400 BAD REQUEST", "BAD REQUEST")
	}
}

func respond(w io.Writer, status string, body string) {
	fmt.Fprintf(w, "HTTP/1.1 %s\r\n", status)
	fmt.Fprintf(w, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(w, "\r\n")
	fmt.Fprintf(w, "%s\r\n", body)
}

func closeConn(conn net.Conn) {
	fmt.Println("Closing connection...")
	conn.Close()
}
