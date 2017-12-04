package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// Try fetchall with longer argument lists, such as samples from the top million web
// sites available at alexa.com. How does the program behave if a web site just doesn’t
// respond? (Section 8.9 describes mechanisms for coping in such cases.)
func main() {
	wg := &sync.WaitGroup{}
	for _, url := range os.Args[1:] {
		go fetch(url, wg)
		wg.Add(1)
	}
	wg.Wait()
}

func fetch(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error while fetching %s: %v\n", url, err)
		return
	}
	size, err := io.Copy(&bytes.Buffer{}, resp.Body)
	if err != nil {
		fmt.Printf("Error while fetching body for %s: %v\n", url, err)
		return
	}
	defer resp.Body.Close()
	time := time.Since(startTime)
	fmt.Printf("%.2fs %d %s\n", time.Seconds(), size, url)
}
