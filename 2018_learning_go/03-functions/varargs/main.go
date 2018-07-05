package main

import "fmt"

func main() {
  vararg()
  vararg(10, 20, 30)
  vararg(11, 22, 33, 44, 55, 66, 77)
}

// Write a function that takes a variable number of ints and print each integer
// on a separate line.
func vararg(xs ...int) {
  for _, x := range xs {
    fmt.Println(x)
  }
}
