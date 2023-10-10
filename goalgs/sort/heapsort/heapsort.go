package heapsort

import (
  "golang.org/x/exp/constraints"

  pq "github.com/bd878/goalgs/ds/priorityqueue"
)

func Heapsort[T constraints.Ordered](a []T, l, r int) {
  q := &pq.SortingTreePQ[T]{}
  q.SetItems(a[l:r])
  q.SetN(r-l)

  for k := pq.Parent(q.N()); k >= 0; k-- {
    q.FixDown(k)
  }

  // sort in place
  for ; !q.Empty(); {
    q.Exchange(l, q.N()-1)
    q.SetN(q.N()-1)
    q.FixDown(l)
  }
}