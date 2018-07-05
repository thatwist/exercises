package main

import "fmt"

func main() {
	sqr := func(n int) int { return n * n }
	fmt.Printf("%v\n", mapi(sqr, []int{}))
	fmt.Printf("%v\n", mapi(sqr, []int{1, 2, 3}))
}

// Write a simple map()-function in Go. It is sufficient for this function only to work for ints.
func mapi(f func(int) int, xs []int) []int {
	nxs := make([]int, len(xs))
	for i, n := range xs {
		nxs[i] = f(n)
	}
	return nxs
}
