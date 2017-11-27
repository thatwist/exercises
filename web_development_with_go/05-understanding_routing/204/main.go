// Building upon the code from the previous problem:
// Extract the code you wrote to READ from the connection using bufio.NewScanner into its own function called "serve".
// Pass the connection of type net.Conn as an argument into this function.
// Add "go" in front of the call to "serve" to enable concurrency and multiple connections.

package main

import (
	"bufio"
	"fmt"
	"io"
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
	io.WriteString(conn, "I see you connected.")
}
