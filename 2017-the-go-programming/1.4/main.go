package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Modify dup2 to print the names of all
// files in which each duplicated line occurs.
func main() {
	counts := make(map[string]int)
	files := make(map[string]map[string]bool)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
			if files[line] == nil {
				files[line] = make(map[string]bool)
			}
			files[line][file] = true
		}
	}

	for line, count := range counts {
		if count > 1 {
			filesForLine := make([]string, 0)
			for f := range files[line] {
				filesForLine = append(filesForLine, f)
			}
			fmt.Println(line, count, strings.Join(filesForLine, ", "))
		}
	}
}
