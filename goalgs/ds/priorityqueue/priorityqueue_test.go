package priorityqueue_test

import (
  "testing"
  "math/rand"
  "sort"
  ds "github.com/bd878/goalgs/ds/priorityqueue"
)

func TestPriorityQueue(t *testing.T) {
  funcs := []func(int) ds.PriorityQueue[int] {
    ds.NewUnorderedArrayPQ[int], ds.NewSortingTreePQ[int],
  }

  size := rand.Intn(10e3)

  for _, fn := range funcs {  
    q := fn(size)
    s := rand.Perm(size)

    for i := 0; i < size; i++ {
      q.Insert(s[i])
    }

    sort.Sort(sort.Reverse(sort.IntSlice(s)))
    for i := 0; !q.Empty(); i++ {
      max := q.GetMax()
      if max != s[i] {
        t.Errorf("=== %dth values not equal : %v != %v\n", i, max, s[i])
      }
    }
  }
}

func TestSortingTree(t *testing.T) {
  for scenario, s := range map[string][]int {
    "random perm": rand.Perm(rand.Intn(10e3)),
    "one-element": []int{5},
    "two elements": []int{2, 1},
  } {
    t.Run(scenario, func(t *testing.T) {
      size := len(s)
      q := ds.NewSortingTreePQ[int](size)

      for i := 0; i < size; i++ {
        q.Insert(s[i])
      }

      if q.Empty() {
        t.Error("sorting tree pq is empty")
      }

      sort.Sort(sort.Reverse(sort.IntSlice(s)))

      for i := 0; i < len(s); i++ {
        max := q.GetMax()
        if max != s[i] {
          t.Errorf("max elements not equal: %v != %v\n", max, s[i])
        }
      }
    })
  }
}