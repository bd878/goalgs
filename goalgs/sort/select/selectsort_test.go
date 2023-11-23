package selectsort_test

import (
  "testing"
  "sort"
  "os"
  "math/rand"
  "flag"

  algs "github.com/bd878/goalgs/sort/select"
  ds "github.com/bd878/goalgs/ds/linkedlist"
)

var size = flag.Int("size", 10000, "size of permutation slice")

func TestMain(m *testing.M) {
  flag.Parse()

  os.Exit(m.Run())
}

func TestSelectsort(t *testing.T) {
  ns := sort.IntSlice(rand.Perm(*size))

  algs.Selectsort(ns)
  if !sort.IsSorted(ns) {
    t.Errorf("slice is not sorted")
  }
}

func TestLLSelectsort(t *testing.T) {
  perm := rand.Perm(10)
  head := ds.InitDumpHeadNode[int]()
  for i := 0; i < len(perm); i++ {
    head.Insert(ds.NewDumpHeadNode[int](perm[i]))
  }

  out := algs.LLSelectsort[int](head)
  result := serializeLL(out)

  if !sort.IsSorted(sort.IntSlice(result)) {
    t.Errorf("not sorted")
  }
}

func serializeLL(head *ds.DumpHeadNode[int]) []int {
  result := make([]int, 0)
  head.Traverse(func(n *ds.DumpHeadNode[int]) {
    result = append(result, n.Item())
  })
  return result
}