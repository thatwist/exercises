package main

import (
	"fmt"
	"net/http"
	"os"
)

// To illustrate the minimum necessary to retrieve information over HTTP,
// here’s a simple program called fetch that fetches the content of each
// specified URL and prints it as uninterpreted text;

// Modify fetch to add the prefix http:// to each argument URL if it is
// missing. You might want to use strings.HasPrefix.

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", url, err)
			continue
		}
		fmt.Printf("%s: %s\n", url, resp.Status)
	}
}
