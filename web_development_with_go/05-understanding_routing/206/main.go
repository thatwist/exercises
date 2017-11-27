// Building upon the code from the previous problem:
// Before we WRITE our RESPONSE , let's WRITE to our RESPONSE the STATUS LINE and some REPONSE HEADERS. Remember the request line and status line:
// REQUEST LINE GET / HTTP/1.1 method SP request-target SP HTTP-version CRLF https://tools.ietf.org/html/rfc7230#section-3.1.1
// RESPONSE (STATUS) LINE HTTP/1.1 302 Found HTTP-version SP status-code SP reason-phrase CRLF https://tools.ietf.org/html/rfc7230#section-3.1.2
// Write the following strings to the response - use io.WriteString for all of the following except the second and third:
// "HTTP/1.1 200 OK\r\n"
// fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
// fmt.Fprint(c, "Content-Type: text/plain\r\n")
// "\r\n"
// Look in your browser "developer tools" under the network tab. Compare the RESPONSE HEADERS from the previous file with the RESPONSE HEADERS in your new solution.

package main

import (
	"bufio"
	"fmt"
	"net"
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
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		fmt.Println(ln)
	}
	body := "Hello world"
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
