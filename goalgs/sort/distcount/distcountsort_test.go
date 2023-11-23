package distcountsort_test

import (
  "testing"
  "sort"

  algs "github.com/bd878/goalgs/sort/distcount"
)

func TestDistcountsort(t *testing.T) {
  a := []int{0, 3, 3, 5, 9, 5, 5, 4, 0, 1, 3, 3, 2}
  algs.Distcount(a, 0, len(a)-1)
  t.Logf("%v\n", a)
  if !sort.IsSorted(sort.IntSlice(a)) {
    t.Error("not sorted")
  }
}