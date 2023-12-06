package qsort_test

import (
  "testing"
  "sort"
  "math/rand"

  algs "github.com/bd878/goalgs/sort/qsort"
)

func TestQSort(t *testing.T) {
  for scenario, alg := range map[string]func([]int, int, int) {
    "qsort recursive": algs.QSortRecursive[int],
    "qsort": algs.QSort[int],
    "qsort insort": algs.QSortInsort[int],
    "hybrid qsort": algs.HybridQSort[int],
    "qsort 3": algs.QSort3[int],
  } {
    t.Run(scenario, func(t *testing.T) {
      perm := rand.Perm(10e3)
      alg(perm, 0, len(perm)-1)

      if !sort.IsSorted(sort.IntSlice(perm)) {
        t.Error("not sorted")
      }
    })
  }
}

func TestFindMedian(t *testing.T) {
  perm := rand.Perm(20)
  k := 5

  sorted := make([]int, len(perm))
  copy(sorted, perm)
  sort.Sort(sort.IntSlice(sorted))

  algs.SelectionMedian[int](perm, 0, len(perm)-1, k)

  expect := sorted[k]

  if perm[k] != expect {
    t.Errorf("%d != %d\n", perm[k], expect)
  }
}

func TestBQSort(t *testing.T) {
  perm := []int{0b100,0b11,0b10,0b1}
  algs.QSortB(perm, 0, len(perm)-1)
  if !sort.IsSorted(sort.IntSlice(perm)) {
    t.Error("not sorted")
  }
}

func TestQSort3Radix(t *testing.T) {
  for scenario, perm := range map[string][]int{
    "two ordered random values": []int{54321, 98845},
    "two reversed random values": []int{98845, 54321},
    "three ordered random values": []int{54321, 98845, 55423},
    "random set": []int{54321, 98845, 55423, 61234, 74389, 18273, 56472},
    "all keys digits same": []int{111, 333, 555, 444, 333, 444, 333, 111},
    "reverse": []int{99, 88, 77, 66, 55, 44, 33, 22},
    "keys with zeroes": []int{54320, 56780, 38029, 64900, 10920, 45607},
  } {
    t.Run(scenario, func(t *testing.T) {
      algs.QSort3Radix(perm, 0, len(perm)-1)
      if !sort.IsSorted(sort.IntSlice(perm)) {
        t.Error("not sorted")
      }
    })
  }
}
