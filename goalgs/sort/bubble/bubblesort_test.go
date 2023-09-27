package main

import (
  "os"
  "sort"
  "flag"
  "math/rand"
  "testing"
)

var size = flag.Int("size", 10e3, "sort size length")

func TestMain(m *testing.M) {
  flag.Parse()

  os.Exit(m.Run())
}

func TestBubbleSort(t *testing.T) {
  ns := sort.IntSlice(rand.Perm(*size))

  Bubblesort(ns)
  if !sort.IsSorted(ns) {
    t.Errorf("perm not sorted")
  }
}