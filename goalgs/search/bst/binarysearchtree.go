package bst

import (
  "fmt"
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
