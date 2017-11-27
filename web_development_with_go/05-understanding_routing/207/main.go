// Building upon the code from the previous problem:
// Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.
// Add this data to your REPONSE so that this data is displayed in the browser.

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
	method := fields[0]
	uri := fields[1]

	body := fmt.Sprintf(`
		Request method: %s
		Request URI: %s`, method, uri)
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
