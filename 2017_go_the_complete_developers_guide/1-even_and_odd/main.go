package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range numbers {
		fmt.Printf("%d is %s\n", i, evenOrOdd(i))
	}
}

func evenOrOdd(i int) string {
	if i%2 == 0 {
		return "even"
	}
	return "odd"
}
