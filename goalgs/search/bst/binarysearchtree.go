package bst

import (
  "fmt"
  "errors"
  "golang.org/x/exp/constraints"
  "github.com/bd878/goalgs/search/types"
)

type BinaryST[K constraints.Ordered, I types.Item[K]] struct {
  head *BTreeNode[I]
}

func NewBinaryST[K constraints.Ordered, I types.Item[K]]() *BinaryST[K, I] {
  return &BinaryST[K, I]{}
}

func (s *BinaryST[K, I]) Init(h *BTreeNode[I]) {
  s.head = h
}

func (s *BinaryST[K, I]) Head() *BTreeNode[I] {
  return s.head
}

func (s *BinaryST[K, I]) Print() {
  s.head.print(func(v *BTreeNode[I], h int) {
    if v == nil {
      fmt.Println("{nil}")
    } else {
      fmt.Printf("%" + fmt.Sprint(h+3) + "v, N: %d\n", v.Item, v.N)
    }
  }, 0)
}

func (s *BinaryST[K, I]) searchR(h *BTreeNode[I], v K) I {
  if h == nil {
    var zero I
    return zero
  }
  k := h.Item.Key()
  if k == v {
    return h.Item
  }
  if v < k {
    return s.searchR(h.L, v)
  } else {
    return s.searchR(h.R, v)
  }
}

func (s *BinaryST[K, I]) insertR(h *BTreeNode[I], x I) {
  h.N += 1

  if x.Key() < h.Item.Key() {
    if h.L == nil {
      h.L = &BTreeNode[I]{Item: x, N: 1}
      return
    }

    s.insertR(h.L, x)
  } else {
    if h.R == nil {
      h.R = &BTreeNode[I]{Item: x, N: 1}
      return
    }

    s.insertR(h.R, x)
  }
}

func (s *BinaryST[K, I]) InsertNonRecursive(x I) {
  v := x.Key()
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  p := s.head
  q := p
  for q != nil {
    q.N += 1

    if v < q.Item.Key() {
      q = q.L
    } else {
      q = q.R
    }

    if q != nil {
      p = q
    }
  }

  if v < p.Item.Key() {
    p.L = &BTreeNode[I]{Item: x, N: 1}
  } else {
    p.R = &BTreeNode[I]{Item: x, N: 1}
  }
}

func (s *BinaryST[K, I]) Search(v K) I {
  return s.searchR(s.head, v)
}

func (s *BinaryST[K, I]) Insert(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  s.insertR(s.head, x)
}

// Rotates tree rights
func rotR[K constraints.Ordered, I types.Item[K]](h **BTreeNode[I]) error {
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

func (s *BinaryST[K, I]) TopRotateR() error {
  if s.head != nil {
    return rotR(&s.head)
  }
  return nil
}

// Rotates tree lefts
func rotL[K constraints.Ordered, I types.Item[K]](h **BTreeNode[I]) error {
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

func (s *BinaryST[K, I]) TopRotateL() error {
  if s.head != nil {
    return rotL(&s.head)
  }
  return nil
}

// insert in root
func (s *BinaryST[K, I]) insertT(h **BTreeNode[I], x I) {
  hv := *h
  hv.N += 1
  if x.Key() < hv.Item.Key() {
    if hv.L == nil {
      hv.L = &BTreeNode[I]{Item: x, N: 1}
    } else {
      s.insertT(&(hv.L), x)
    }
    if err := rotR(h); err != nil {
      panic(err)
    }
  } else {
    if hv.R == nil {
      hv.R = &BTreeNode[I]{Item: x, N: 1}
    } else {
      s.insertT(&(hv.R), x)
    }
    if err := rotL(h); err != nil {
      panic(err)
    }
  }
}

func (s *BinaryST[K, I]) InsertInRoot(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  s.insertT(&s.head, x)
}

func (s *BinaryST[K, I]) Sort() {
  panic("not implemented")
}

func (s *BinaryST[K, I]) selectR(h *BTreeNode[I], k int) I {
  if h == nil {
    var zero I
    return zero
  }
  var t int
  if h.L != nil {
    t = h.L.N
  }
  if t > k {
    return s.selectR(h.L, k)
  }
  if t < k {
    // t elements on left branch,
    // we need (t-k)'th smallest
    // element from the right branch
    return s.selectR(h.R, k-t-1)
  }
  return h.Item
}

func (s *BinaryST[K, I]) Select(k int) I {
  return s.selectR(s.head, k)
}

func (s *BinaryST[K, I]) Count() int {
  return s.Head().N
}

// Put k'th least element in the root.
// Rotates when necessary.
// 0'th is the smallest element
func (s *BinaryST[K, I]) partitionR(h **BTreeNode[I], k int) {
  hv := *h
  var t int
  if hv.L != nil {
    t = hv.L.N
  }
  if t > k {
    s.partitionR(&(hv.L), k)
    rotR(h)
  }
  if t < k {
    s.partitionR((&hv.R), k-t-1)
    rotL(h)
  }
}

func (s *BinaryST[K, I]) Partition(k int) {
  s.partitionR(&s.head, k)
}

func (s *BinaryST[K, I]) joinLR(a *BTreeNode[I], b *BTreeNode[I]) *BTreeNode[I] {
  if b == nil {
    return a
  }
  s.partitionR(&b, 0)
  b.L = a
  b.N += a.N
  return b
}

// N updates in joinLR.
// see also randomRemoveR in optimizations
func (s *BinaryST[K, I]) removeR(h **BTreeNode[I], v K) {
  if (*h) == nil {
    return
  }
  hv := *h
  w := hv.Item.Key()
  if v < w {
    if hv.L != nil {
      s.removeR(&(hv.L), v)
    }
  }
  if v > w {
    if hv.R != nil {
      s.removeR(&(hv.R), v)
    }
  }
  if v == w {
    // write new root in parent,
    // erase previous item
    *h = s.joinLR(hv.L, hv.R)
  }
}

func (s *BinaryST[K, I]) Remove(x I) {
  s.removeR(&(s.head), x.Key())
}

func (s *BinaryST[K, I]) joinR(a *BTreeNode[I], b *BTreeNode[I]) *BTreeNode[I] {
  if b == nil {
    return a
  }
  if a == nil {
    return b
  } 
  s.insertT(&a, b.Item)
  s.head = a
  a.L = s.joinR(a.L, b.L)
  a.R = s.joinR(a.R, b.R)
  a.N = 1
  if a.L != nil {
    a.N += a.L.N
  }
  if a.R != nil {
    a.N += a.R.N
  }
  return a
}

func (s *BinaryST[K, I]) Join(b *BinaryST[K, I]) {
  s.head = s.joinR(s.head, b.Head())
}

func (s *BinaryST[K, I]) heightR(h *BTreeNode[I]) int {
  if h.L == nil && h.R == nil {
    return 0
  }

  var lHeight, rHeight int
  if h.L != nil {
    lHeight = s.heightR(h.L)
  }
  if h.R != nil {
    rHeight = s.heightR(h.R)
  }

  return max(lHeight, rHeight) + 1
}

func (s *BinaryST[K, I]) Height() int {
  return s.heightR(s.head)
}