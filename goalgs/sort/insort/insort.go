package insort

import (
  "sort"
  "golang.org/x/exp/constraints"
)

func Insort(nums sort.Interface) {
  if nums.Len() > 1 {
    for i := 0; i < nums.Len(); i++ {
      for j := i; j > 0 && nums.Less(j, j - 1); j-- {
        nums.Swap(j, j - 1)
      }
    }
  }
}

func InsortRange[T constraints.Ordered](a []T, l, r int) {
  if len(a) > 1 {
    for i := 0; i < len(a); i++ {
      for j := i; j > 0 && a[i] < a[j-1]; j-- {
        a[i], a[j-1] = a[j-1], a[i]
      }
    }
  }
}
