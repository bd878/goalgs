package pqsort_test

import (
  "testing"
  "math/rand"
  "sort"

  alg "github.com/bd878/goalgs/sort/pqsort"
)

func TestPQSort(t *testing.T) {
  size := rand.Intn(10e3)
  perm := rand.Perm(size)

  alg.PQSort[int](perm, 0, len(perm)-1)
  if !sort.IsSorted(sort.IntSlice(perm)) {
    t.Errorf("not sorted")
  }
}