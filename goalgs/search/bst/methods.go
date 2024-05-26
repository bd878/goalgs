package bst

import (
  "errors"
  "math/rand"
  "golang.org/x/exp/constraints"
  "github.com/bd878/goalgs/search/types"
)

func heightR[I interface{}](h *BTreeNode[I]) int {
  if h.L == nil && h.R == nil {
    return 0
  }

  var lHeight, rHeight int
  if h.L != nil {
    lHeight = heightR(h.L)
  }
  if h.R != nil {
    rHeight = heightR(h.R)
  }

  return max(lHeight, rHeight) + 1
}

func joinR[K constraints.Ordered, I types.Item[K]](a *BTreeNode[I], b *BTreeNode[I]) *BTreeNode[I] {
  if b == nil {
    return a
  }
  if a == nil {
    return b
  } 
  insertT(&a, b.Item)

  a.L = joinR(a.L, b.L)
  a.R = joinR(a.R, b.R)
  a.N = 1
  if a.L != nil {
    a.N += a.L.N
  }
  if a.R != nil {
    a.N += a.R.N
  }
  return a
}

func joinLR[I interface{}](a *BTreeNode[I], b *BTreeNode[I]) *BTreeNode[I] {
  if b == nil {
    return a
  }
  partitionR(&b, 0)
  b.L = a
  b.N += a.N
  return b
}

// N updates in joinLR.
// see also randomRemoveR in optimizations
func removeR[K constraints.Ordered, I types.Item[K]](h **BTreeNode[I], v K) {
  if (*h) == nil {
    return
  }
  hv := *h
  w := hv.Item.Key()
  if v < w {
    if hv.L != nil {
      removeR(&(hv.L), v)
    }
  }
  if v > w {
    if hv.R != nil {
      removeR(&(hv.R), v)
    }
  }
  if v == w {
    // write new root in parent,
    // erase previous item
    *h = joinLR(hv.L, hv.R)
  }
}

// Put k'th least element in the root.
// Rotates when necessary.
// 0'th is the smallest element
func partitionR[I interface{}](h **BTreeNode[I], k int) {
  hv := *h
  var t int
  if hv.L != nil {
    t = hv.L.N
  }
  if t > k {
    partitionR(&(hv.L), k)
    rotR(h)
  }
  if t < k {
    partitionR((&hv.R), k-t-1)
    rotL(h)
  }
}

func selectR[I interface{}](h *BTreeNode[I], k int) I {
  if h == nil {
    var zero I
    return zero
  }
  var t int
  if h.L != nil {
    t = h.L.N
  }
  if t > k {
    return selectR(h.L, k)
  }
  if t < k {
    // t elements on left branch,
    // we need (t-k)'th smallest
    // element from the right branch
    return selectR(h.R, k-t-1)
  }
  return h.Item
}

// insert in root
func insertT[K constraints.Ordered, I types.Item[K]](h **BTreeNode[I], x I) {
  hv := *h
  hv.N += 1
  if x.Key() < hv.Item.Key() {
    if hv.L == nil {
      hv.L = &BTreeNode[I]{Item: x, N: 1}
    } else {
      insertT(&(hv.L), x)
    }
    if err := rotR(h); err != nil {
      panic(err)
    }
  } else {
    if hv.R == nil {
      hv.R = &BTreeNode[I]{Item: x, N: 1}
    } else {
      insertT(&(hv.R), x)
    }
    if err := rotL(h); err != nil {
      panic(err)
    }
  }
}

// Rotates tree lefts
func rotL[I interface{}](h **BTreeNode[I]) error {
  if *h == nil {
    return errors.New("h is nil")
  }
  if (*h).R == nil {
    return errors.New("h.R is nil, nothing to left-rotate")
  }

  x := (*h).R

  // x is right node of h. x is the new header
  var xLn, xRn, hLn int
  if x.R != nil {
    xRn = x.R.N
  }
  if x.L != nil {
    xLn = x.L.N
  }
  if (*h).L != nil {
    hLn = (*h).L.N
  }

  // prevh is previous root
  prevh := *h
  // *h now points to new root, it is previous right branch
  *h = (*h).R
  // save left branch of previous right branch (2-level-nesting)
  prevL := (*h).L
  // new root left branch now points on previous root
  (*h).L = prevh
  // previous root right branch now points
  // on left branch of previous right branch 
  prevh.R = prevL

  // +1 this node
  prevh.N = 1 + hLn + xLn
  (*h).N = 1 + xRn + prevh.N

  return nil
}

// Rotates tree rights
func rotR[I interface{}](h **BTreeNode[I]) error {
  if *h == nil {
    return errors.New("h is nil")
  }
  if (*h).L == nil {
    return errors.New("h.L is nil, nothing to right-rotate")
  }

  x := (*h).L

  // x is right node of h. x is the new header
  var xLn, xRn, hRn int
  if x.R != nil {
    xRn = x.R.N
  }
  if x.L != nil {
    xLn = x.L.N
  }
  if (*h).R != nil {
    hRn = (*h).R.N
  }

  prevh := *h
  *h = (*h).L
  prevR := (*h).R
  (*h).R = prevh
  prevh.L = prevR

  prevh.N = 1 + xRn + hRn
  (*h).N = 1 + xLn + prevh.N

  return nil
}

func searchR[K constraints.Ordered, I types.Item[K]](h *BTreeNode[I], v K) I {
  if h == nil {
    var zero I
    return zero
  }
  k := h.Item.Key()
  if k == v {
    return h.Item
  }
  if v < k {
    return searchR(h.L, v)
  } else {
    return searchR(h.R, v)
  }
}

func insertR[K constraints.Ordered, I types.Item[K]](h *BTreeNode[I], x I) {
  h.N += 1

  if x.Key() < h.Item.Key() {
    if h.L == nil {
      h.L = &BTreeNode[I]{Item: x, N: 1}
      return
    }

    insertR(h.L, x)
  } else {
    if h.R == nil {
      h.R = &BTreeNode[I]{Item: x, N: 1}
      return
    }

    insertR(h.R, x)
  }
}

func randomInsertR[K constraints.Ordered, I types.Item[K]](h **BTreeNode[I], x I) {
  hv := *h
  if rand.Intn(RAND_MAX) < int(RAND_MAX/(hv.N+1)) {
    insertT(h, x)
    return
  }

  hv.N += 1
  if x.Key() < hv.Item.Key() {
    if hv.L == nil {
      hv.L = &BTreeNode[I]{Item: x, N: 1}
      return
    }
    randomInsertR(&(hv.L), x)
  } else {
    if hv.R == nil {
      hv.R = &BTreeNode[I]{Item: x, N: 1}
      return
    }
    randomInsertR(&(hv.R), x)
  }
}

// must be used in RandomRemove instead of joinLR.
// we suppose that a and b are two branches of one tree,
// so a < b, and hence a.L < a.R < b.L < b.R
// Public for tests
func randomJoinLR[I interface{}](a *BTreeNode[I], b *BTreeNode[I]) *BTreeNode[I] {
  if a == nil {
    return b
  }
  if b == nil {
    return a
  }

  if rand.Intn(RAND_MAX) /
    (RAND_MAX /
      (a.N + b.N) + 1) < a.N {
    a.R = randomJoinLR(a.R, b) // b != nil, a.R ?= nil

    a.N = 1 + a.R.N
    if a.L != nil {
      a.N += a.L.N
    }

    return a
  }
  b.L = randomJoinLR(a, b.L) // a != nil, b.L ?= nil

  b.N = 1 + b.L.N
  if b.R != nil {
    b.N += b.R.N
  }

  return b
}

// copied from removeR
// return true if key found and value removed
func randomRemoveR[K constraints.Ordered, I types.Item[K]](h **BTreeNode[I], v K) bool {
  if (*h) == nil {
    return false
  }
  hv := *h
  w := hv.Item.Key()
  if v < w {
    if hv.L != nil {
      if randomRemoveR(&(hv.L), v) {
        hv.N -= 1
      }
    }
  }
  if v > w {
    if hv.R != nil {
      if randomRemoveR(&(hv.R), v) {
        hv.N -= 1
      }
    }
  }
  if v == w {
    // write new root in parent,
    // erase previous item
    *h = randomJoinLR(hv.L, hv.R)

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

func splay[K constraints.Ordered, I types.Item[K]](h **BTreeNode[I], x I) {
  hv := *h
  w := hv.Item.Key()

  hv.N += 1
  if x.Key() < w { // left branch, right-right or left-right
    if hv.L == nil {
      hv.L = &BTreeNode[I]{Item: x, N: 1}
      return
    }

    if x.Key() < hv.L.Item.Key() {
      splay(&(hv.L), x) // insert left
      rotR(h) // rotate right from root
    } else {
      splay(&(hv.R), x) // insert right
      rotL(&(hv.L)) // rotate left from child node
    }
    rotR(h) // rotate right
  } else { // right branch, left-left or right-left
    if hv.R == nil {
      hv.R = &BTreeNode[I]{Item: x, N: 1}
      return
    }

    if x.Key() > hv.R.Item.Key() {
      splay(&(hv.R), x) // insert right
      rotL(h) // rotate left from root
    } else {
      splay(&(hv.L), x) // insert left
      rotR(&(hv.R)) // rotate right from child node
    }
    rotL(h)
  }
}
