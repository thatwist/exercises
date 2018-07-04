package main

import "fmt"

func main() {
	main1()
	main2()
	main3()
}

// Create a loop with the for construct. Make it loop 10 times and print out
// the loop counter with the fmt package.
func main1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// Rewrite the loop from 1 to use goto. The keyword for may not be used.
func main2() {
	var i int
Loop:
	fmt.Println(i)
	i++
	if i < 10 {
		goto Loop
	}
}

// Rewrite the loop again so that it fills an array and then prints that array
// to the screen.
func main3() {
	var x [10]int
	for i, _ := range x {
		x[i] = i
	}
	fmt.Printf("%v\n", x)
}
