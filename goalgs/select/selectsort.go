package main

import "sort"

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