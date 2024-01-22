package linell

import (
  "fmt"
  "golang.org/x/exp/constraints"
  ll "github.com/bd878/goalgs/ds/linkedlist"
  "github.com/bd878/goalgs/search/types"
)

type OrderedST[K constraints.Ordered, I types.Item[K]] struct {
  head ll.LLNode[I]
  n int
  zero I
}

func NewOrderedST[K constraints.Ordered, I types.Item[K]]() *OrderedST[K, I] {
  head := ll.InitPtrLL[I]()
  var zero I
  return &OrderedST[K, I]{head: head, zero: zero, n: 0}
}

func (s *OrderedST[K, I]) Count() int {
  return s.n
}

func (s *OrderedST[K, I]) Search(key K) I {
  res := s.zero
  s.head.Traverse(func(n ll.LLNode[I]) {
    if n.Item().Key() == key {
      res = n.Item()
    }
  })
  return res
}

func (s *OrderedST[K, I]) Insert(x I) {
  p := s.head
  s.head.Traverse(func(n ll.LLNode[I]) {
    if n.Item().Key() < x.Key() {
      p = n
    }
  })
  if p == s.head && !p.IsEmpty() {
    if x.Key() < p.Item().Key() {
      s.head = ll.NewPtrNode[I](x)
      s.head.SetNext(p)
    } else {
      p.Insert(ll.NewPtrNode[I](x))
    }
  } else {
    p.Insert(ll.NewPtrNode[I](x))
  }
  s.n += 1
}

func (s *OrderedST[K, I]) Remove(x I) {
  p := s.head
  found := false
  s.head.Traverse(func(n ll.LLNode[I]) {
    if found { return; }
    if n.Item().Key() == x.Key() {
      found = true
    } else {
      p = n
    }
  })
  if found {
    if p == s.head {
      s.head = p.Next()
    } else {
      p.DeleteNext()
    }
    s.n -= 1
  }
}

func (s *OrderedST[K, I]) Select(i int) I {
  res := s.zero
  j := 0
  s.head.Traverse(func(n ll.LLNode[I]) {
    if j == i {
      res = n.Item()
    }
    j += 1
  })
  return res
}

func (s *OrderedST[K, I]) Sort() {
  /* sorted on Insert */
}

func (s *OrderedST[K, I]) Print() {
  s.head.Traverse(func(n ll.LLNode[I]) {
    fmt.Printf("%v ", n.Item().Key())
  })
  fmt.Println()
}