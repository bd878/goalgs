package tournament

import (
  "golang.org/x/exp/constraints"

  tree "github.com/bd878/goalgs/ds/tree"
)

func Max[T constraints.Ordered](a []T, l, r int) *tree.BTreeNode[T] {
  m := (l+r)/2
  x := tree.NewNode[T](a[m])
  if l == r {
    return x
  }

  x.L = Max(a, l, m)
  x.R = Max(a, m+1, r)
  lv, rv := x.L.V, x.R.V
  if lv > rv {
    x.V = lv
  } else {
    x.V = rv
  }

  return x
}