package linearray

import (
  "fmt"
  "golang.org/x/exp/constraints"
  "github.com/bd878/goalgs/search/types"
)

const MAX_N int = 10

type OrderedST[K constraints.Ordered, I types.Item[K]] struct {
  n int
  St []I
  zero I
}

func NewOrderedST[K constraints.Ordered, I types.Item[K]]() *OrderedST[K, I] {
  st := make([]I, MAX_N)
  var zero I
  return &OrderedST[K, I]{n:0, St:st, zero: zero}
}

func (s *OrderedST[K, I]) Count() int {
  return s.n
}

// search is consecutive intentionally
func (s *OrderedST[K, I]) Search(v K) I {
  var i int
  for ; i < s.n; i++ {
    if s.St[i].Key() >= v { break; }
  }
  if s.St[i].Key() == v { return s.St[i]; }
  return s.zero
}

func (s *OrderedST[K, I]) Sort() { /* sorted on insert */ }

func (s *OrderedST[K, I]) Insert(x I) {
  v := x.Key()
  i := s.n
  // right shift elements
  for i > 0 && v < s.St[i-1].Key() {
    s.St[i] = s.St[i-1]
    i -= 1
  }
  s.St[i] = x
  s.n += 1
}

func (s *OrderedST[K, I]) Remove(x I) {
  v := x.Key()
  i := 0
  for ; i < s.n && v != s.St[i].Key(); i++ {}
  if s.St[i].Key() != v { return; }

  j := s.n-1
  t := s.St[j]
  for j != i {
    t, s.St[j-1] = s.St[j-1], t
    j -= 1
  }
  s.St[s.n-1] = s.zero
  s.n -= 1
}

func (s *OrderedST[K, I]) print() {
  for i := 0; i < s.n; i++ {
    fmt.Print(s.St[i].Key(), " ")
  }
  fmt.Println()
}

func (s *OrderedST[K, I]) Select(i int) I {
  if i < s.n {
    return s.St[i]
  }
  return s.zero
}
