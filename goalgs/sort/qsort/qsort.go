package qsort

import (
  "golang.org/x/exp/constraints"

  ds "github.com/bd878/goalgs/ds/stack"
)

func QSortRecursive[T constraints.Ordered](a []T, l, r int) {
  if l >= r {
    return;
  }
  i := Part[T](a, l, r)
  QSortRecursive[T](a, l, i-1)
  QSortRecursive[T](a, i+1, r)
}

func Part[T constraints.Ordered](a []T, l, r int) int {
  v := a[r]
  i, j := l, r-1
  for i < j {
    for ; a[i] < v && i < j; i++ {}
    for ; a[j] > v && j > i; j-- {}
    if i < j {
      a[i], a[j] = a[j], a[i]
    }
  }
  if a[r] < a[i] {
    a[i], a[r] = a[r], a[i]
    return i
  } else {
    return r
  }
}

func push(s ds.Stack[int], l, r int) {
  s.Push(r)
  s.Push(l)
}

func QSort[T constraints.Ordered](a []T, l, r int) {
  s := ds.NewArrStack[int]()
  push(s, l, r)

  for !s.IsEmpty() {
    l, _ = s.Pop()
    r, _ = s.Pop()
    if r <= l {
      continue
    }

    i := Part[T](a, l, r)
    if i > r-i { // i is above the half
      push(s, l, i-1) // larger period goes first
      push(s, i+1, r) // take smaller distance on next iteration
    } else {
      push(s, i+1, r) // larger period goes first
      push(s, l, i-1) // take smaller then
    }
  }
}