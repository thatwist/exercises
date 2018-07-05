package main

import "fmt"

func main() {
	p := plusTwo()
	fmt.Printf("%v\n", p(2))
	p2 := plusX(2)
	fmt.Printf("%v\n", p2(2))
}

// 1. Write a function that returns a function that performs a +2 on integers.
// Name the function plusTwo. You should then be able do the following:
func plusTwo() func(int) int {
	return func(n int) int {
		return n + 2
	}
}

// 2. Generalize the function from above and create a plusX(x) which returns
// functions that add x to an integer.
func plusX(x int) func(int) int {
	return func(n int) int {
		return n + x
	}
}
