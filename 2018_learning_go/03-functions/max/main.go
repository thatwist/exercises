package main

import "fmt"

func main() {
	fmt.Printf("%d\n", max([]int{}))
	fmt.Printf("%d\n", max([]int{1}))
	fmt.Printf("%d\n", max([]int{1, 2}))
	fmt.Printf("%d\n", max([]int{10, 1}))
	fmt.Printf("%d\n", max([]int{10, 30, 20}))
	fmt.Printf("%d\n", max([]int{10, 30, 20, 30}))
}

// Write a function that finds the maximum value in an int slice ([]int).
func max(xs []int) int {
	var m int
	for _, n := range xs {
		if m < n {
			m = n
		}
	}
	return m
}
