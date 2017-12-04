package main

import (
	"fmt"
	"os"
)

// Modify the echo program to also print os.Args[0],
// the name of the command that invoked it.
// (echo prints all the passed arguments to it)
func main() {
	fmt.Println(os.Args)
}
