package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		if isValid(words) {
			count++
		}
	}
	fmt.Println(count)
}

func isValid(words []string) bool {
	if len(words) < 2 {
		return true
	}
	for i, w1 := range words[:len(words)-1] {
		for _, w2 := range words[i+1:] {
			if w1 == w2 {
				return false
			}
		}
	}
	return true
}
