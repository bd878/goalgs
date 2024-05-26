package bst

import (
  "fmt"
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
  return searchR(s.head, v)
}

func (s *BinaryST[K, I]) Insert(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  insertR(s.head, x)
}

func (s *BinaryST[K, I]) TopRotateR() error {
  if s.head != nil {
    return rotR(&s.head)
  }
  return nil
}

func (s *BinaryST[K, I]) TopRotateL() error {
  if s.head != nil {
    return rotL(&s.head)
  }
  return nil
}

func (s *BinaryST[K, I]) InsertInRoot(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  insertT(&s.head, x)
}

func (s *BinaryST[K, I]) Sort() {
  panic("not implemented")
}

func (s *BinaryST[K, I]) Select(k int) I {
  return selectR(s.head, k)
}

func (s *BinaryST[K, I]) Count() int {
  return s.Head().N
}

func (s *BinaryST[K, I]) Partition(k int) {
  partitionR(&s.head, k)
}

func (s *BinaryST[K, I]) Remove(x I) {
  removeR(&(s.head), x.Key())
}

func (s *BinaryST[K, I]) Join(b *BinaryST[K, I]) {
  s.head = joinR(s.head, b.Head())
}

func (s *BinaryST[K, I]) Height() int {
  return heightR(s.head)
}