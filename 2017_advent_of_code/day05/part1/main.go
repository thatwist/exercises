package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := loadNumbers(scanner)
	var cursor, steps int
	for cursor < len(numbers) {
		steps++
		jump := numbers[cursor]
		numbers[cursor]++
		cursor += jump
	}
	fmt.Println(steps)
}

func loadNumbers(scanner *bufio.Scanner) []int {
	var numbers []int
	for scanner.Scan() {
		ln := scanner.Text()
		i, err := strconv.Atoi(ln)
		if err == nil {
			numbers = append(numbers, i)
		}
	}
	return numbers
}
