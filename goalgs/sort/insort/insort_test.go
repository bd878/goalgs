package insort_test

import (
  "sort"
  "math/rand"
  "testing"

  algs "github.com/bd878/goalgs/sort/insort"
)

func TestInsort(t *testing.T) {
  ns := sort.IntSlice(rand.Perm(10e2))

  algs.Insort(ns)
  if !sort.IsSorted(ns) {
    t.Error("slice is not sorted")
  }
}
