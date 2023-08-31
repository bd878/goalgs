package main

import "fmt"
import "math/rand"

func main() {
  ns := rand.Perm(10)

  Insort(ns)
  println(ns)
}

func Insort(nums []int) {
  if len(nums) > 1 {
    for j := 1; j < len(nums); j++ {
      key := nums[j]
      for i := j - 1; i >= 0 && nums[i] >= key; i-- {
        nums[i + 1] = nums[i]
      }
    }
  }
}

func println(nums []int) {
  for _, n := range nums {
    fmt.Printf("%d ", n)
  }
  fmt.Println()
}
