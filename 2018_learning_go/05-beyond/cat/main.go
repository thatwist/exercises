package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "os"
)

// Write a program which mimics the Unix program cat.
// Make it support the -n flag, where each line is numbered.
func main() {
  n := flag.Bool("n", false, "Number the output lines, starting at 1.")
  flag.Parse()
  // for simplification stdin only
  scanner := bufio.NewScanner(os.Stdin)
  var i int
  for scanner.Scan() {
    i++
    if *n {
      io.WriteString(os.Stdout, fmt.Sprintf("%6d\t%s\n", i, scanner.Text()))
    } else {
      io.WriteString(os.Stdout, fmt.Sprintf("%s\n", scanner.Text()))
    }
  }
}
