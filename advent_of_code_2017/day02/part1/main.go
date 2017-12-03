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
	var sum int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := strings.Fields(line)
		first, _ := strconv.Atoi(numbers[0])
		min := first
		max := first
		for _, f := range numbers {
			n, err := strconv.Atoi(f)
			if err != nil {
				continue
			}
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
		sum += (max - min)
	}
	fmt.Println(sum)
}
