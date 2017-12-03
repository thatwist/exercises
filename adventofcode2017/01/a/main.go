package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

// http://adventofcode.com/2017/day/1
func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	var sum int
	// assume last one is new line char so rewrite it with first element
	bytes[len(bytes)-1] = bytes[0]
	for i, digitByte := range bytes[:len(bytes)-1] {
		if digitByte == bytes[i+1] {
			digit, _ := strconv.Atoi(string(digitByte))
			sum += digit
		}
	}
	fmt.Println(sum)
}
