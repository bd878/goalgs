package qsort

import (
  "golang.org/x/exp/constraints"
)

func QSort[T constraints.Ordered](a []T, l, r int) {
  if l >= r {
    return;
  }
  i := Part[T](a, l, r)
  QSort[T](a, l, i-1)
  QSort[T](a, i+1, r)
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