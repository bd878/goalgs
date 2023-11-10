package binaryqueue

import (
  "golang.org/x/exp/constraints"

  btree "github.com/bd878/goalgs/ds/tree"
)

type Handle[T constraints.Ordered] *btree.BTreeNode[T]

type BinaryPQ[T constraints.Ordered] struct {
  bq []*btree.BTreeNode[T]
}

func NewBinaryPQ[T constraints.Ordered](size int) *BinaryPQ[T] {
  return &BinaryPQ[T]{
    bq: make([]*btree.BTreeNode[T], size),
  }
}

func (q *BinaryPQ[T]) Bq() []*btree.BTreeNode[T] {
  return q.bq
}

func (q *BinaryPQ[T]) IsEmpty() bool {
  for _, v := range q.bq {
    if !v.IsEmpty() {
      return false
    }
  }
  return true
}

func (q *BinaryPQ[T]) Insert(item T) Handle[T] {
  t := btree.NewNode[T](item)
  c := t
  for i := 0; i < len(q.bq) && c != nil; i++ {
    if q.bq[i] == nil {
      q.bq[i] = c
      c = nil
    } else {
      c = Pair(c, q.bq[i])
      q.bq[i] = nil
    }
  }

  if c != nil {
    q.bq = append(q.bq, c)
  }

  return Handle[T](t)
}

func (q *BinaryPQ[T]) GetMax() T {
  var v T
  var maxi int = -1

  for i := 0; i < len(q.bq); i++ {
    if q.bq[i] != nil {
      if maxi == -1 || v < q.bq[i].V {
        maxi = i
        v = q.bq[i].V
      }
    }
  }

  x := q.bq[maxi].L
  // size of btree is ln(count elements) - number of bytes in binary number
  temp := NewBinaryPQ[T](maxi)

  for i := maxi; i > 0; i-- {
    temp.bq[i-1] = x
    x = x.R
    temp.bq[i-1].R = nil
  }

  q.bq[maxi] = nil
  q.Join(temp)

  return v
}

// Joins tree p into q, suppose that q size > p size
func (q *BinaryPQ[T]) Join(p *BinaryPQ[T]) {
  var c *btree.BTreeNode[T]

  // implied len(p.bq) < len(q.bq)
  for i := 0; i < len(p.bq); i++ {
    // c, b, a
    switch valuesToBinary(c != nil, p.bq[i] != nil, q.bq[i] != nil) {
    case 2: // 010
      q.bq[i] = p.bq[i]
    case 3: // 011
      c = Pair[T](q.bq[i], p.bq[i])
      q.bq[i] = nil
    case 4: // 100
      q.bq[i] = c
      c = nil
    case 5: // 101
      c = Pair[T](c, q.bq[i])
      q.bq[i] = nil
    case 6: // 110
    case 7: // 111
      c = Pair[T](c, p.bq[i])
    }
  }

  if c != nil {
    q.bq = append(q.bq, c)
  }
}

func valuesToBinary(c, b, a bool) int {
  res := 0
  if c { res += 4 }
  if b { res += 2 }
  if a { res += 1 }
  return res
}