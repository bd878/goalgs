package bst

import (
  "math/rand"
)

const RAND_MAX int = 10e3

func (s *BinaryST[K, I]) randomInsertR(h **BTreeNode[I], x I) {
  hv := *h
  if rand.Intn(RAND_MAX) < int(RAND_MAX/(hv.N+1)) {
    s.insertT(h, x)
    return
  }

  hv.N += 1
  if x.Key() < hv.Item.Key() {
    if hv.L == nil {
      hv.L = &BTreeNode[I]{Item: x, N: 1}
      return
    }
    s.randomInsertR(&(hv.L), x)
  } else {
    if hv.R == nil {
      hv.R = &BTreeNode[I]{Item: x, N: 1}
      return
    }
    s.randomInsertR(&(hv.R), x)
  }
}

func (s *BinaryST[K, I]) RandomInsert(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  s.randomInsertR(&s.head, x)
}