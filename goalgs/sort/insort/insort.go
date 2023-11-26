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
  if r-l > 0 {
    for i := l; i <= r; i++ {
      for j := i; j > 0 && a[j] < a[j-1]; j-- {
        a[j], a[j-1] = a[j-1], a[j]
      }
    }
  }
}
