package binaryqueue

import (
  "golang.org/x/exp/constraints"
  btree "github.com/bd878/goalgs/ds/tree"
)

// Joins two equal size binary sorting trees
// and returns the greates node of two
func Pair[T constraints.Ordered](p, q *btree.BTreeNode[T]) *btree.BTreeNode[T] {
  if p.IsEmpty() {
    return q
  } else if q.IsEmpty() {
    return p
  }

  var result *btree.BTreeNode[T]

  if (p.V < q.V) {
    p.R = q.L
    q.L = p // q.R is empty
    result = q
  } else {
    q.R = p.L
    p.L = q // n.R is empty
    result = p
  }

  return result
}