package bst

import (
  "fmt"
  "errors"
  "golang.org/x/exp/constraints"
  "github.com/bd878/goalgs/search/types"
)

// TODO: rewrite on ds/tree/binarytree
type BTreeNode[I interface{}] struct {
  Item I
  L *BTreeNode[I]
  R *BTreeNode[I]
}

func (n *BTreeNode[I]) print(printer func(I, int), h int) {
  if n == nil {
    var zero I
    printer(zero, h)
    return;
  }

  printer(n.Item, h)
  n.R.print(printer, h+1)
  n.L.print(printer, h+1)
}

type BinaryST[K constraints.Ordered, I types.Item[K]] struct {
  head *BTreeNode[I]
}

func NewBinaryST[K constraints.Ordered, I types.Item[K]]() *BinaryST[K, I] {
  return &BinaryST[K, I]{}
}

func (s *BinaryST[K, I]) Head() *BTreeNode[I] {
  return s.head
}

func (s *BinaryST[K, I]) Print() {
  s.head.print(func(v I, h int) {
    fmt.Printf("%" + fmt.Sprint(h+3) + "v\n", v)
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
  if x.Key() < h.Item.Key() {
    if h.L == nil {
      h.L = &BTreeNode[I]{Item: x}
      return
    }

    s.insertR(h.L, x)
  } else {
    if h.R == nil {
      h.R = &BTreeNode[I]{Item: x}
      return
    }

    s.insertR(h.R, x)
  }
}

func (s *BinaryST[K, I]) InsertNonRecursive(x I) {
  v := x.Key()
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x}
    return
  }
  p := s.head
  q := p
  for q != nil {
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
    p.L = &BTreeNode[I]{Item: x}
  } else {
    p.R = &BTreeNode[I]{Item: x}
  }
}

func (s *BinaryST[K, I]) Search(v K) I {
  return s.searchR(s.head, v)
}

func (s *BinaryST[K, I]) Insert(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x}
    return
  }
  s.insertR(s.head, x)
}

// Rotates tree rights
func (s *BinaryST[K, I]) rotR(h **BTreeNode[I]) error {
  if *h == nil {
    return errors.New("h is nil")
  }
  if (*h).L == nil {
    return errors.New("h.L is nil, nothing to right-rotate")
  }
  // h is a pointer to current header link
  // x is a left node of h, x is the new header
  x := (*h).L
  (*h).L = x.R
  // h must become right branch of x
  x.R = *h
  *h = x

  return nil
}

func (s *BinaryST[K, I]) TopRotateR() error {
  if s.head != nil {
    return s.rotR(&s.head)
  }
  return nil
}

// Rotates tree lefts
func (s *BinaryST[K, I]) rotL(h **BTreeNode[I]) error {
  if *h == nil {
    return errors.New("h is nil")
  }
  if (*h).R == nil {
    return errors.New("h.R is nil, nothing to left-rotate")
  }
  // x is right node of h. x is the new header
  x := (*h).R
  (*h).R = x.L
  x.L = *h
  *h = x

  return nil
}

func (s *BinaryST[K, I]) TopRotateL() error {
  if s.head != nil {
    return s.rotL(&s.head)
  }
  return nil
}

func (s *BinaryST[K, I]) insertT(h **BTreeNode[I], x I) {
  hv := *h
  if x.Key() < hv.Item.Key() {
    if hv.L == nil {
      hv.L = &BTreeNode[I]{Item: x}
      return
    }
    s.insertT(&(hv.L), x)
    if err := s.rotR(h); err != nil {
      panic(err)
    }
  } else {
    if hv.R == nil {
      hv.R = &BTreeNode[I]{Item: x}
      return
    }
    s.insertT(&(hv.R), x)
    if err := s.rotL(h); err != nil {
      panic(err)
    }
  }
}

func (s *BinaryST[K, I]) InsertInRoot(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x}
    return
  }
  s.insertT(&s.head, x)
}

func (s *BinaryST[K, I]) Sort() {
  panic("not implemented")
}

func (s *BinaryST[K, I]) Remove(x I) {
  panic("not implemented")
}

func (s *BinaryST[K, I]) Select(_ int) I {
  panic("not implemented")
}

func (s *BinaryST[K, I]) Count() int {
  panic("not implemented")
}
