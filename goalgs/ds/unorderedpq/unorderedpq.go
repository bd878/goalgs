package unorderedpq

import (
  "golang.org/x/exp/constraints"
  dll "github.com/bd878/goalgs/ds/doublelinkedlist"
)

type UnorderedPQ[T constraints.Ordered] struct {
  l *dll.List[T]
}

// no reserve memory
func New[T constraints.Ordered]() *UnorderedPQ[T] {
  return &UnorderedPQ[T]{
    l: dll.New[T](),
  }
}

func (q *UnorderedPQ[T]) Empty() bool {
  return q.l.Empty()
}

// fast
func (q *UnorderedPQ[T]) Insert(item T) *dll.Node[T] {
  return q.l.Insert(dll.NewNode[T](item))
}

// slow
func (q *UnorderedPQ[T]) GetMax() T {
  var max T
  if q.l.Empty() {
    return max 
  }

  x := q.l.Head()
  for i := x; i != q.l.Tail().Next(); i = i.Next() {
    if i.Item() > x.Item() {
      x = i
    }
  }
  max = x.Item()
  q.l.Delete(x)
  return max
}

// fast
func (q *UnorderedPQ[T]) Change(n *dll.Node[T], item T) {
  n.SetItem(item)
}

// fast
// func (q *UnorderedPQ[T]) Join(p *UnorederedPQ[T]) {}

func (q *UnorderedPQ[T]) Remove(n *dll.Node[T]) {
  q.l.Delete(n)
}
