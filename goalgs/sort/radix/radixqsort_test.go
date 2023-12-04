package radix_test

import (
  "testing"
  "sort"
  algs "github.com/bd878/goalgs/sort/radix"
)

func reverse(n int) int {
  var result, digit int
  for n != 0 {
    digit = n % 10
    n = n/10 
    result = result*10 + digit
  }
  return result
}

func TestRadixMSD(t *testing.T) {
  for scenario, perm := range map[string][]int{
    "random set": []int{54321, 98845, 55423, 61234, 74389, 18273, 56472},
    "all keys digits same": []int{111, 333, 555, 444, 333, 444, 333, 111},
    "reverse": []int{99, 88, 77, 66, 55, 44, 33, 22},
    "keys with zeroes": []int{54320, 56780, 38029, 64900, 10920, 45607},
  } {
    t.Run(scenario, func(t *testing.T) {
      algs.RadixMSD(perm, 0, len(perm)-1)

      if !sort.IsSorted(sort.IntSlice(perm)) {
        t.Errorf("not sorted")
      }
    })
  }
}
