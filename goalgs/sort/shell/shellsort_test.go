package shellsort_test

import (
  "testing"
  "math/rand"
  "sort"

  algs "github.com/bd878/goalgs/sort/shell"
)

func TestShellSort(t *testing.T) {
  perm := rand.Perm(10e3)

  algs.Shellsort[int](perm, 0, len(perm))
  if !sort.IsSorted(sort.IntSlice(perm)) {
    t.Error("not sorted")
  }
}