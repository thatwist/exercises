package main

import "fmt"

func main() {
	fmt.Printf("%v\n", bubblesort([]int{}))
	fmt.Printf("%v\n", bubblesort([]int{2}))
	fmt.Printf("%v\n", bubblesort([]int{3, 2}))
	fmt.Printf("%v\n", bubblesort([]int{3, 2, 1}))
	fmt.Printf("%v\n", bubblesort([]int{1, 2, 3, 4}))
	fmt.Printf("%v\n", bubblesort([]int{1, 3, 2, 4}))
	fmt.Printf("%v\n", bubblesort([]int{1, 3, 2, 1, 5}))
}

// Write a function that performs a bubble sort on a slice of ints.
func bubblesort(xs []int) []int {
	size := len(xs)
	sxs := make([]int, size)
	copy(sxs, xs)
	if size > 1 {
		sorted := false
		for !sorted {
			sorted = true
			for i := 0; i < size-1; i++ {
				if sxs[i] > sxs[i+1] {
					tmp := sxs[i]
					sxs[i] = sxs[i+1]
					sxs[i+1] = tmp
					sorted = false
				}
			}
		}
	}
	return sxs
}
