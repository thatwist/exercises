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
      o1, o2 := pop2(stack)
      stack.Push(strconv.Itoa(o2 + o1))
    case "-":
      o1, o2 := pop2(stack)
      stack.Push(strconv.Itoa(o2 - o1))
    case "*":
      o1, o2 := pop2(stack)
      stack.Push(strconv.Itoa(o2 * o1))
    case "/":
      o1, o2 := pop2(stack)
      stack.Push(strconv.Itoa(o2 / o1))
    default:
      stack.Push(token)
    }
  }
  result, _ := strconv.Atoi(stack.Pop())
  return result
}

func pop2(stack *data.Stack) (i1 int, i2 int) {
  i1, _ = strconv.Atoi(stack.Pop())
  i2, _ = strconv.Atoi(stack.Pop())
  return
}
