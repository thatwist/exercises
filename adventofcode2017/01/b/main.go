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
	half := (len(bytes) - 1) / 2
	for i, digitByte := range bytes[:half] {
		if digitByte == bytes[i+half] {
			digit, _ := strconv.Atoi(string(digitByte))
			sum += digit * 2
		}
	}
	fmt.Println(sum)
}
