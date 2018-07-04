package main

import "fmt"

func main() {
	fmt.Printf("%f\n", avg([]float64{0.1, 0.3, 0.5, 0.8}))
	fmt.Printf("%f\n", avg([]float64{}))
}

// Write code to calculate the average of a float64 slice. In a later exercise
// you will make it into a function.
func avg(xs []float64) float64 {
	size := len(xs)
	if size == 0 {
		return 0.0
	}
	var sum float64
	for _, v := range xs {
		sum += v
	}
	return sum / float64(size)
}
