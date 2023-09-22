package main

import (
  "testing"
  "math/rand"
  "os"
  "flag"
  "sort"
)

const MAX_N = 10e1

var (
  alen = flag.Int("alen", 10e3, "a length")
  blen = flag.Int("blen", 10e2, "b length")
)

func TestMain(m *testing.M) {
  flag.Parse()

  os.Exit(m.Run())
}

func TestMergeAB(t *testing.T) {
  a, b := getSorted()

  c := MergeAB(a, a.Len(), b, b.Len())
  if !sort.IsSorted(sort.IntSlice(c)) {
    t.Error("c is not sorted")
  }
}

func TestMerge(t *testing.T) {
  a, b := getSorted()

  c := make([]int, 0, a.Len()+b.Len())
  c = append(c, a...)
  c = append(c, b...)

  Merge(c, 0, a.Len(), a.Len()+b.Len())

  if !sort.IsSorted(sort.IntSlice(c)) {
    t.Error("c is not sorted")
  }
}

func getSorted() (a, b sort.IntSlice) {
  a = make([]int, *alen)
  b = make([]int, *blen)
  for i := 0; i < *alen; i++ {
    a[i] = rand.Intn(MAX_N)
  }
  for i := 0; i < *blen; i++ {
    b[i] = rand.Intn(MAX_N)
  }

  sort.Sort(a)
  sort.Sort(b)

  return a, b
}