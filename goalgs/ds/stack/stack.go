package stack

import "errors"

type Stack[T interface{}] struct {
  values []T
  top int
}

func New[T interface{}]() *Stack[T] {
  return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
  s.values = append(s.values, v)
  s.top += 1
}

func (s *Stack[T]) Pop() (T, error) {
  if s.IsEmpty() {
    var d T
    return d, errors.New("stack is empty")
  }

  s.top -= 1
  val := s.values[s.top]
  s.values = s.values[:s.top]

  return val, nil
 }

func (s *Stack[T]) IsEmpty() bool {
  return s.top == 0
}
