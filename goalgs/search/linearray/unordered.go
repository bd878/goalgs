package linearray

import (
  "golang.org/x/exp/constraints"
  "github.com/bd878/goalgs/search/types"
)

type UnorderedST[K constraints.Ordered, I types.Item[K]] struct {
  n int
  st []I
  zero I
}

func NewUnorderedST[K constraints.Ordered, I types.Item[K]]() *UnorderedST[K, I] {
  st := make([]I, MAX_N)
  var zero I
  return &UnorderedST[K, I]{n:0, st:st, zero: zero}
}

func (s *UnorderedST[K, I]) Count() int {
  return s.n
}

func (s *UnorderedST[K, I]) Insert(x I) {
  s.st[s.n] = x
  s.n += 1
}

func (s *UnorderedST[K, I]) Remove(x I) {
  for i, v := range s.st {
    if v.Key() == x.Key() {
      s.st[s.n-1], s.st[i] = s.st[i], s.st[s.n-1]
      s.n -= 1

      return
    }
  }
}

func (s *UnorderedST[K, I]) Search(p K) I {
  for _, v := range s.st {
    if v.Key() == p {
      return v
    }
  }
  return s.zero
}

func (s *UnorderedST[K, I]) Select(i int) I {
  /* ST is unordered */
  panic("TODO: not implemented")
}

func (s *UnorderedST[K, I]) Sort() {
  /* ST is unordered */
}