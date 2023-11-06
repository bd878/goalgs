package stack

import "errors"

type ArrStack[T interface{}] struct {
  values []T
  top int
}

func NewArrStack[T interface{}]() Stack[T] {
  return &ArrStack[T]{}
}

func (s *ArrStack[T]) Push(v T) {
  s.values = append(s.values, v)
  s.top += 1
}

func (s *ArrStack[T]) Pop() (T, error) {
  if s.IsEmpty() {
    var d T
    return d, errors.New("stack is empty")
  }

  s.top -= 1
  val := s.values[s.top]
  s.values = s.values[:s.top]

  return val, nil
 }

func (s *ArrStack[T]) IsEmpty() bool {
  return s.top == 0
}
