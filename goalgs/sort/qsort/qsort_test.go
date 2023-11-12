package qsort_test

import (
  "testing"
  "sort"
  "math/rand"

  mysort "github.com/bd878/goalgs/sort/qsort"
)

func TestQSort(t *testing.T) {
  perm := rand.Perm(10e3)
  mysort.QSort[int](perm, 0, len(perm)-1)

  if !sort.IsSorted(sort.IntSlice(perm)) {
    t.Error("not sorted")
  }
}