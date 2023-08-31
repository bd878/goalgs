package main

import "math/rand"
import "testing"
import "flag"

var size = flag.Int("size", 10000, "size of permutation slice")

func BenchmarkInsort(b *testing.B) {
  flag.Parse()

  for i := 0; i < b.N; i++ {
    r := rand.New(rand.NewSource(int64(i)))
    ns := r.Perm(*size)

    Insort(ns)
    if !isSorted(ns) {
      b.Errorf("slice is not sorted")
    }
  }
}

func isSorted(nums []int) bool {
  for i := 0; i < len(nums) - 1; i++ {
    if nums[i] > nums[i+1] {
      return false
    }
  }
  return true
}