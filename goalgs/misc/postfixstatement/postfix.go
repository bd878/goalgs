package postfixstatement

import (
  stack "github.com/bd878/goalgs/ds/stack"
)

func Evaluate(expression string) int {
  s := stack.New[int]()

  var a, b int
  for _, v := range expression {
    if v == '+' {
      a, _ = s.Pop()
      b, _ = s.Pop()
      s.Push(a + b)
    } else if v == '*' {
      a, _ = s.Pop()
      b, _ = s.Pop()
      s.Push(a * b)
    } else if v >= '0' && v <= '9' {
      // transform ASCII symbols to digit numbers
      s.Push(int(v - '0'))
    }
  }

  res, _ := s.Pop()
  return res
}