package data

// Stack is a basic implemention of a data structure Stack which organizes
// collection of data (strings in this case) in a LIFO (Last In First out) manner.
// The implementation is not thread safe.
type Stack struct {
  container []string
}

func (s *Stack) Size() int {
  return len(s.container)
}

func (s *Stack) Push(str string) {
  s.container = append(s.container, str)
}

func (s *Stack) Pop() string {
  str := s.container[s.Size()-1]
  s.container = s.container[:s.Size()-1]
  return str
}
