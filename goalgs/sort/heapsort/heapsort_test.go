package heapsort_test

import (
  "testing"
  "math/rand"
  "sort"

  alg "github.com/bd878/goalgs/sort/heapsort"
)

func TestHeapsort(t *testing.T) {
  size := rand.Intn(10e3)
  perm := rand.Perm(size)

  alg.Heapsort[int](perm, 0, size)

  if !sort.IsSorted(sort.IntSlice(perm)) {
    t.Errorf("not sorted")
  }
}
