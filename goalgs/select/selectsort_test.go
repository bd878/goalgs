package main

import (
  "testing"
  "sort"
  "math/rand"
  "flag"
)

var size = flag.Int("size", 10000, "size of permutation slice")

func TestMain(m *testing.M) {
  flag.Parse()

  os.Exit(m.Run())
}

func TestSelectsort(t *testing.T) {
  ns := sort.IntSlice(rand.Perm(*size))

  Selectsort(ns)
  if !sort.IsSorted(ns) {
    t.Errorf("slice is not sorted")
  }
}