package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// To illustrate the minimum necessary to retrieve information over HTTP,
// here’s a simple program called fetch that fetches the content of each
// specified URL and prints it as uninterpreted text;

// The function call io.Copy(dst, src) reads from src and writes to dst. Use it
// instead of ioutil.ReadAll to copy the response body to os.Stdout without
// requiring a buffer large enough to hold the entire stream. Be sure to check
// the error result of io.Copy.
func main() {
	url := os.Args[1]
	if url == "" {
		fmt.Println("URL is missing")
		return
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while fetching %s: %v\n", url, err)
		return
	}
	i, err := io.Copy(os.Stdout, res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading response of %s: %v\n", url, err)
		return
	}
	fmt.Printf("Written %d bytes\n", i)
}
