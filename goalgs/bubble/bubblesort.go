package main

import "sort"

func Bubblesort(a sort.Interface) {
  for i := 0; i < a.Len(); i++ {
    for j := a.Len()-1; j > 0; j-- {
      if a.Less(j, j-1) {
        a.Swap(j, j-1)
      }
    }
  }
}