package pqsort

import (
  "golang.org/x/exp/constraints"
  pq "github.com/bd878/goalgs/ds/priorityqueue"
)

func PQSort[T constraints.Ordered](a []T, l, r int) {
  q := pq.NewSortingTreePQ[T](r-l+1)
  for _, v := range a {
    q.Insert(v)
  }
  for k := r; k >= l; k-- {
    a[k] = q.GetMax()
  }
}