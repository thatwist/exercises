package main

import "fmt"

func main() {
	fmt.Printf("%v\n", fib(2))
	fmt.Printf("%v\n", fib(3))
	fmt.Printf("%v\n", fib(4))
	fmt.Printf("%v\n", fib(10))
}

// Write a function that takes an int value and gives that many terms of the
// Fibonacci sequence.
func fib(n int) []int {
	xs := make([]int, n)
	xs[0], xs[1] = 1, 1
	for i := 2; i < n; i++ {
		xs[i] = xs[i-1] + xs[i-2]
	}
	return xs
}
