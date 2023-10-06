package priorityqueue_test

import (
  "testing"
  "math/rand"
  "sort"
  ds "github.com/bd878/goalgs/ds/priorityqueue"
)

func TestPriorityQueue(t *testing.T) {
  size := rand.Intn(10e3)

  s := rand.Perm(size)
  q := ds.NewUnorderedArrayPQ[int](size)
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