package linkedlist

import (
  "golang.org/x/exp/constraints"
)

type DumpHeadNode[T constraints.Ordered] struct {
  value T
  next *DumpHeadNode[T]
}

func InitDumpHeadNode[T constraints.Ordered]() *DumpHeadNode[T] {
  return &DumpHeadNode[T]{}
}

func NewDumpHeadNode[T constraints.Ordered](value T) *DumpHeadNode[T] {
  return &DumpHeadNode[T]{value: value}
}

func (x *DumpHeadNode[T]) Insert(t *DumpHeadNode[T]) *DumpHeadNode[T] {
  if x.Next() != nil {
    t.SetNext(x.Next())
  }
  return x.SetNext(t)
}

func (x *DumpHeadNode[T]) DeleteNext() *DumpHeadNode[T] {
  if x.Next() != nil {
    result := x.Next()
    x.SetNext(x.Next().Next())
    return result
  }
  return nil
}

func (x *DumpHeadNode[T]) SetNext(t *DumpHeadNode[T]) *DumpHeadNode[T] {
  x.next = t
  return t
}

func (x *DumpHeadNode[T]) Next() *DumpHeadNode[T] {
  return x.next
}

func (x *DumpHeadNode[T]) Item() T {
  return x.value
}

func (x *DumpHeadNode[T]) Traverse(fn func(*DumpHeadNode[T])) {
  for t := x.Next(); t != nil; t = t.Next() {
    fn(t)
  }
}
