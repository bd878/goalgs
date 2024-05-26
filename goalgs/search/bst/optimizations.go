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

// must be used in RandomRemove instead of joinLR.
// we suppose that a and b are two branches of one tree,
// so a < b, and hence a.L < a.R < b.L < b.R
// Public for tests
func (s *BinaryST[K, I]) RandomJoinLR(a *BTreeNode[I], b *BTreeNode[I]) *BTreeNode[I] {
  if a == nil {
    return b
  }
  if b == nil {
    return a
  }
  if rand.Intn(RAND_MAX) /
    (RAND_MAX /
      (a.N + b.N) + 1) < a.N {
    a.R = s.RandomJoinLR(a.R, b) // b != nil, a.R ?= nil

    a.N = 1 + a.R.N
    if a.L != nil {
      a.N += a.L.N
    }

    return a
  } else {
    b.L = s.RandomJoinLR(a, b.L) // a != nil, b.L ?= nil

    b.N = 1 + b.L.N
    if b.R != nil {
      b.N += b.R.N
    }

    return b
  }
}

// copied from removeR
// return true if key found and value removed
func (s *BinaryST[K, I]) randomRemoveR(h **BTreeNode[I], v K) bool {
  if (*h) == nil {
    return false
  }
  hv := *h
  w := hv.Item.Key()
  if v < w {
    if hv.L != nil {
      if s.randomRemoveR(&(hv.L), v) {
        hv.N -= 1
      }
    }
  }
  if v > w {
    if hv.R != nil {
      if s.randomRemoveR(&(hv.R), v) {
        hv.N -= 1
      }
    }
  }
  if v == w {
    // write new root in parent,
    // erase previous item
    *h = s.RandomJoinLR(hv.L, hv.R)

    hv = *h
    hv.N = 1
    if hv.L != nil {
      hv.N += hv.L.N
    }
    if hv.R != nil {
      hv.N += hv.R.N
    }
    return true
  }
  return false
}

// break interface, remove by key, not by item
func (s *BinaryST[K, I]) RandomRemove(v K) {
  s.randomRemoveR(&(s.head), v)
}

func (s *BinaryST[K, I]) RandomInsert(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  s.randomInsertR(&s.head, x)
}

// We assume, that s and b are already well-balanced
func (s *BinaryST[K, I]) RandomJoin(b *BinaryST[K, I]) {
  n := s.head.N
  if rand.Intn(RAND_MAX) /
    (RAND_MAX /
      (n + b.Head().N) + 1) < n {
    s.head = s.joinR(b.Head(), s.head)
  } else {
    s.head = s.joinR(s.head, b.Head())
  }
}

func (s *BinaryST[K, I]) splay(h **BTreeNode[I], x I) {
  hv := *h
  w := hv.Item.Key()

  hv.N += 1
  if x.Key() < w { // left branch, right-right or left-right
    if hv.L == nil {
      hv.L = &BTreeNode[I]{Item: x, N: 1}
      return
    }

    if x.Key() < hv.L.Item.Key() {
      s.splay(&(hv.L), x) // insert left
      rotR(h) // rotate right from root
    } else {
      s.splay(&(hv.R), x) // insert right
      rotL(&(hv.L)) // rotate left from child node
    }
    rotR(h) // rotate right
  } else { // right branch, left-left or right-left
    if hv.R == nil {
      hv.R = &BTreeNode[I]{Item: x, N: 1}
      return
    }

    if x.Key() > hv.R.Item.Key() {
      s.splay(&(hv.R), x) // insert right
      rotL(h) // rotate left from root
    } else {
      s.splay(&(hv.L), x) // insert left
      rotR(&(hv.R)) // rotate right from child node
    }
    rotL(h)
  }
}

func (s *BinaryST[K, I]) SplayInsert(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  s.splay(&s.head, x)
}