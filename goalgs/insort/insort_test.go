package main

import (
  "sort"
  "math/rand"
  "testing"
  "flag"
)

var size = flag.Int("size", 10000, "size of permutation slice")

func BenchmarkInsort(b *testing.B) {
  flag.Parse()

  for i := 0; i < b.N; i++ {
    r := rand.New(rand.NewSource(int64(i)))
    ns := sort.IntSlice(r.Perm(*size))

    Insort(ns)
    if !sort.IsSorted(ns) {
      b.Errorf("slice is not sorted")
    }
  }
}
