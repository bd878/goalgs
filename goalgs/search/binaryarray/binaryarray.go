package binaryarray

import (
  "golang.org/x/exp/constraints"
  "github.com/bd878/goalgs/search/types"
  "github.com/bd878/goalgs/search/linearray"
)

const MAX_N int = 10

type BinaryST[K constraints.Ordered, I types.Item[K]] struct {
  linearray.OrderedST[K, I]
}

func NewBinaryST[K constraints.Ordered, I types.Item[K]]() *BinaryST[K, I] {
  a := BinaryST[K, I]{
    *(linearray.NewOrderedST[K, I]()),
  }
  return &a
}

func (s *BinaryST[K, I]) searchR(l, r int, v K) I {
  var zero I
  if l > r {
    return zero
  }

  m := (l+r)/2
  if v == s.St[m].Key() {
    return s.St[m]
  }
  if l == r {
    return zero
  }

  if v < s.St[m].Key() {
    return s.searchR(l, m-1, v)
  } else {
    return s.searchR(m+1, r, v)
  }
}

func (s *BinaryST[K, I]) Search(v K) I {
  return s.searchR(0, s.Count()-1, v)
}