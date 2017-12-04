package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Provide a filename as an argument")
	}
	fileName := os.Args[1]
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatal("Error", err)
	}
	io.Copy(os.Stdout, file)
}
