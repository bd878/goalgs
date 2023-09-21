package main

import (
  "testing"
  "math/rand"
  "os"
  "flag"
  "sort"
)

var (
  alen = flag.Int("a", 10e3, "a length")
  blen = flag.Int("b", 10e2, "b length")
)

func TestMain(m *testing.M) {
  flag.Parse()

  os.Exit(m.Run())
}

func TestMergeAB(t *testing.T) {
  a := sort.IntSlice(rand.Perm(*alen))
  b := sort.IntSlice(rand.Perm(*blen))

  sort.Sort(a)
  sort.Sort(b)

  c := MergeAB(a, a.Len(), b, b.Len())
  if !sort.IsSorted(sort.IntSlice(c)) {
    t.Error("c is not sorted")
  }
}