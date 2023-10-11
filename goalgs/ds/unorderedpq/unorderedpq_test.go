package unorderedpq_test

import (
  "testing"
  "math/rand"
  "sort"

  pq "github.com/bd878/goalgs/ds/unorderedpq"
)

func TestUnorderedPQ(t *testing.T) {
  size := rand.Intn(10e3)
  perm := rand.Perm(size)

  q := pq.New[int]()
  for _, v := range perm {
    q.Insert(v)
  }

  if q.Empty() {
    t.Errorf("=== queue is declared empty")
  }

  sort.Sort(sort.Reverse(sort.IntSlice(perm)))
  for i, v := range perm {
    max := q.GetMax()
    if max != v {
      t.Errorf("=== %dth elements not equal: %v != %v\n", i, max, v)
    }
  }
}