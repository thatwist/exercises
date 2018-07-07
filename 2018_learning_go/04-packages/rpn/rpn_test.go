package rpn

import "testing"

func TestEval(t *testing.T) {
  examples := []struct{
    in string
    out int
  }{
    {in: "", out: 0},
    {in: "2 3 +", out: 5},
    {in: "7 2 3 + *", out: 35},
    {in: "7 2 3 + * 2 * 7 / 3 -", out: 7},
  }

  for _, example := range examples {
    r := Eval(example.in)
    if r != example.out {
      t.Logf("'%s' should evaluate to %d, but was %d", example.in, example.out, r)
      t.Fail()
    }
  }
}

