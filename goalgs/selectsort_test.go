package main

import "testing"
import "sort"
import "math/rand"
import "flag"

var size = flag.Int("size", 10000, "size of permutation slice")

func TestSelectsort(t *testing.T) {
	flag.Parse()

  ns := sort.IntSlice(rand.Perm(*size))

  Selectsort(ns)
  if !IsSorted(ns) {
    t.Errorf("slice is not sorted")
  }
}