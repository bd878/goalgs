package main

import "sort"
import "fmt"

func Println(nums []int) {
  for _, n := range nums {
    fmt.Printf("%d ", n)
  }
  fmt.Println()
}

func IsSorted(nums sort.Interface) bool {
  for i := 0; i < nums.Len() - 1; i++ {
    if nums.Less(i + 1, i) {
      return false
    }
  }
  return true
}
