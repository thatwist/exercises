package data

import "testing"

func TestSize(t *testing.T) {
  stack := &Stack{}
  if stack.Size() != 0 {
    t.Log("Initial Stack size should be 0")
    t.Fail()
  }
  stack.Push("foo")
  if stack.Size() != 1 {
    t.Log("Stack size should be 1 after pushing one element")
    t.Fail()
  }
  stack.Pop()
  if stack.Size() != 0 {
    t.Log("Stack size should be 0 after push followed by pop")
    t.Fail()
  }
}

func TestPushPop(t *testing.T) {
  stack := &Stack{}
  stack.Push("foo")
  stack.Push("bar")
  if stack.Pop() != "bar" {
    t.Log("Pop should remove element that was pushed last")
    t.Fail()
  }
  if stack.Pop() != "foo" {
    t.Log("Second pop should remove element pushed before last")
    t.Fail()
  }
}
