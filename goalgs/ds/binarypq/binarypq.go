package binaryqueue

import (
  "golang.org/x/exp/constraints"
)

type Handle[T constraints.Ordered] *Node[T]

type BinaryPQ[T constraints.Ordered] struct {
  bq []*Node[T]
}

func NewBinaryPQ[T constraints.Ordered](size int) *BinaryPQ[T] {
  return &BinaryPQ[T]{
    bq: make([]*Node[T], size),
  }
}

func (q *BinaryPQ[T]) IsEmpty() bool {
  return false
}

func (q *BinaryPQ[T]) Insert(item T) Handle[T] {
  t := NewNode[T](item)
  c := t
  for i := 0; i < len(q.bq) && c != nil; i++ {
    if q.bq[i] == nil {
      q.bq[i] = c
      c = nil
    } else {
      c = c.Pair(q.bq[i])
      q.bq[i] = nil
    }
  }
  return Handle[T](t)
}

func (q *BinaryPQ[T]) GetMax() T {
  panic("not implemented")
}

func (q *BinaryPQ[T]) Change(h Handle[T], item T) {}

func (q *BinaryPQ[T]) Join(p *BinaryPQ[T]) {}
