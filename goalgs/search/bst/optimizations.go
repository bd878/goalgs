package bst

import (
  "math/rand"
)

const RAND_MAX int = 10e3

func (s *BinaryST[K, I]) RandomJoinLR(a *BTreeNode[I], b *BTreeNode[I]) *BTreeNode[I] {
  return randomJoinLR(a, b)
}

// break interface, remove by key, not by item
func (s *BinaryST[K, I]) RandomRemove(v K) {
  randomRemoveR(&(s.head), v)
}

func (s *BinaryST[K, I]) RandomInsert(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  randomInsertR(&s.head, x)
}

// We assume, that s and b are already well-balanced
func (s *BinaryST[K, I]) RandomJoin(b *BinaryST[K, I]) {
  n := s.head.N
  if rand.Intn(RAND_MAX) /
    (RAND_MAX /
      (n + b.Head().N) + 1) < n {
    s.head = joinR(b.Head(), s.head)
  } else {
    s.head = joinR(s.head, b.Head())
  }
}

func (s *BinaryST[K, I]) SplayInsert(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  splay(&s.head, x)
}