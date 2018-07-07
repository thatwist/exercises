package rpn

import "testing"

func TestEmpty(t *testing.T) {
  if Eval("") != 0 {
    t.Log("Empty expression should be evaluted to 0")
    t.Fail()
  }
}

func TestEvalSum(t *testing.T) {
  if Eval("2 3 +") != 5 {
    t.Log("2 3 + should be evaluted to 5")
    t.Fail()
  }
}

func TestEvalSumMultiply(t *testing.T) {
  if Eval("7 2 3 + *") != 35 {
    t.Log("7 2 3 + should be evaluted to 35")
    t.Fail()
  }
}
