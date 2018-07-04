package main

import "fmt"

// Write a program that prints the numbers from 1 to 100.
// But for multiples of three print, "Fizz" instead of the number,
// and for multiples of five, print "Buzz".
// For numbers which are multiples of both three and five, print "FizzBuzz".
func main() {
  for i := 1; i <= 100; i++ {
    m3 := i % 3 == 0
    m5 := i % 5 == 0
    switch {
    case m3 && m5:
      fmt.Println("FizzBuzz")
    case m3:
      fmt.Println("Fizz")
    case m5:
      fmt.Println("Buzz")
    default:
      fmt.Printf("%d\n", i)
    }
  }
}
