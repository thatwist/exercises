package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum, d int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := stringsToInts(strings.Fields(line))
		for i, x := range numbers[:len(numbers)-1] {
			for _, y := range numbers[i+1:] {
				if x%y == 0 {
					d = x / y
					continue
				} else if y%x == 0 {
					d = y / x
					continue
				}
			}
		}
		sum += d
	}
	fmt.Println(sum)
}

func stringsToInts(strings []string) []int {
	ints := make([]int, len(strings))
	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}
