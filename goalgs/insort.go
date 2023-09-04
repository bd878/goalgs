package main

import "sort"
import "math/rand"

func main() {
  ns := sort.IntSlice(rand.Perm(10))

  Insort(ns)
  Println(ns)
}

func Insort(nums sort.Interface) {
  if nums.Len() > 1 {
    for i := 0; i < nums.Len(); i++ {
      for j := i; j > 0 && nums.Less(j, j - 1); j-- {
        nums.Swap(j, j - 1)
      }
    }
  }
}
