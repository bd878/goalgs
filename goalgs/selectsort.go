package main

import "sort"
import "math/rand"

func main() {
  ns := sort.IntSlice(rand.Perm(10))

  Selectsort(ns)
  Println(ns)
}

func Selectsort(nums sort.Interface) {
  for i := 0; i < nums.Len(); i++ {
    min := i
    for j := i; j < nums.Len(); j++ {
      if nums.Less(j, min) {
        min = j
      }
    }
    nums.Swap(min, i)
  }
}