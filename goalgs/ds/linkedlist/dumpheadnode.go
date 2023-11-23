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
  if x.next != nil {
    t.next = x.next
  }
  x.next = t
  return t
}

func (x *DumpHeadNode[T]) DeleteNext() *DumpHeadNode[T] {
  if x.next != nil {
    result := x.next
    x.next = result.next
    result.next = nil
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

func (x *DumpHeadNode[T]) IsEmpty() bool {
  var empty T
  return x.value == empty && x.next == nil
}

func (x *DumpHeadNode[T]) Traverse(fn func(*DumpHeadNode[T])) {
  for t := x.next; t != nil; t = t.next {
    fn(t)
  }
}
