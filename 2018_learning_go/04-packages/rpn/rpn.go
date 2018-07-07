package rpn

import (
  "strings"
  "strconv"
  "github.com/dejan/exercises/2018_learning_go/04-packages/stack/data"
)

// Eval evalutes expression written in RPN (Reverse Polish Notation)
func Eval(str string) int {
  stack := &data.Stack{}
  tokens := strings.Split(str, " ")
  for _, token := range tokens {
    switch token {
    case "+":
      o1, _ := strconv.Atoi(stack.Pop())
      o2, _ := strconv.Atoi(stack.Pop())
      r := o2 + o1
      stack.Push(strconv.Itoa(r))
    case "-":
      o1, _ := strconv.Atoi(stack.Pop())
      o2, _ := strconv.Atoi(stack.Pop())
      r := o2 - o1
      stack.Push(strconv.Itoa(r))
    case "*":
      o1, _ := strconv.Atoi(stack.Pop())
      o2, _ := strconv.Atoi(stack.Pop())
      r := o2 * o1
      stack.Push(strconv.Itoa(r))
    case "/":
      o1, _ := strconv.Atoi(stack.Pop())
      o2, _ := strconv.Atoi(stack.Pop())
      r := o2 / o1
      stack.Push(strconv.Itoa(r))
    default:
      stack.Push(token)
    }
  }
  result, _ := strconv.Atoi(stack.Pop())
  return result
}
