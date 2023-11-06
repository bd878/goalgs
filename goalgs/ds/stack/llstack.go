package stack

import (
  "errors"

  ll "github.com/bd878/goalgs/ds/linkedlist"
)

type LLStack[T interface{}] struct {
  head *ll.PtrNode[T]
}

func NewLLStack[T interface{}]() Stack[T] {
  return &LLStack[T]{head: ll.InitPtrLL[T]()}
}

func (s *LLStack[T]) Push(v T) {
  s.head.Insert(ll.NewPtrNode[T](v))
}

func (s *LLStack[T]) Pop() (T, error) {
  if s.IsEmpty() {
    var d T
    return d, errors.New("stack is empty")
  }
  v := s.head.DeleteNext()
  return v.Item(), nil
}

func (s *LLStack[T]) IsEmpty() bool {
  return s.head.IsEmpty()
}