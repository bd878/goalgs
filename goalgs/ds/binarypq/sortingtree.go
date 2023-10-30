package binaryqueue

import (
  "golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
  V T
  L *Node[T]
  R *Node[T]
}

func NewNode[T constraints.Ordered](v T) *Node[T] {
  return &Node[T]{V: v}
}

func (n *Node[T]) SetL(l *Node[T]) {
  n.L = l
}

func (n *Node[T]) SetR(r *Node[T]) {
  n.R = r
}

// returns greater node
func (n *Node[T]) Pair(q *Node[T]) *Node[T] {
  var result *Node[T]

  if (n.V < q.V) {
    n.SetR(q.L)
    q.SetL(n) // q.R is empty
    result = q
  } else {
    q.SetR(n.L)
    n.SetL(q) // n.R is empty
    result = n
  }

  return result
}