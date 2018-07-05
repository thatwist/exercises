package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := &stack{}
	fmt.Printf("%v\n", s)
	s.push(10)
	s.push(20)
	fmt.Printf("%v\n", s)
	fmt.Printf("%d\n", s.pop())
	fmt.Printf("%v\n", s)
	s.push(30)
	fmt.Printf("%d\n", s.pop())
	fmt.Printf("%d\n", s.pop())
}

// Create a simple stack which can hold a fixed number of ints. It does not have
// to grow beyond this limit. Define push -- put something on the stack -- and
// pop -- retrieve something from the stack -- functions. The stack should be a
// LIFO (last in, first out) stack.
type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(x int) {
	if s.i > 9 {
		return
	}
	s.data[s.i] = x
	s.i++
}

func (s *stack) pop() int {
	s.i--
	x := s.data[s.i]
	s.data[s.i] = 0
	return x
}

func (s stack) String() string {
	str := make([]string, s.i)
	for i := 0; i < s.i; i++ {
		str[i] = strconv.Itoa(s.data[i])
	}
	return "[ " + strings.Join(str, ", ") + " ]"
}
