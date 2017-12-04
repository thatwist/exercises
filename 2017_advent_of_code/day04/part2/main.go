package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
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
			if (w1 == w2) || anagrams(w1, w2) {
				return false
			}
		}
	}
	return true
}

func anagrams(s1, s2 string) bool {
	a1 := strings.Split(s1, "")
	a2 := strings.Split(s2, "")
	sort.Strings(a1)
	sort.Strings(a2)
	return reflect.DeepEqual(a1, a2)
}
