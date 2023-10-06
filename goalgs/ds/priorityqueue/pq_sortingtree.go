package priorityqueue

import (
  "golang.org/x/exp/constraints"
)

type SortingTreePQ[T constraints.Ordered] struct {
  pq []T
  n int
}

func NewSortingTreePQ[T constraints.Ordered](maxn int) PriorityQueue[T] {
  return &SortingTreePQ[T]{
    pq: make([]T, maxn+1),
  }
}

func (q *SortingTreePQ[T]) Empty() bool {
  return q.n == 0
}

func (q *SortingTreePQ[T]) Insert(item T) {
  q.n += 1
  q.pq[q.n] = item
  q.fixUp(q.n)
}

func (q *SortingTreePQ[T]) GetMax() T {
  max := q.pq[1]
  // max is last
  q.exchange(1, q.n)
  // balance tree without last
  q.n -= 1
  q.fixDown(1)

  return max
}

func (q *SortingTreePQ[T]) fixUp(k int) {
  // parent is smaller than child
  for ; k > 1 && q.pq[k/2] < q.pq[k]; {
    q.exchange(k, k/2)
    k = k/2
  }
}

func (q *SortingTreePQ[T]) fixDown(k int) {
  for ; 2*k <= q.n; {
    j := 2*k
    if j < q.n && q.pq[j] < q.pq[j+1] {
      // other child
      j += 1
    }
    // child is larger than parent
    if q.pq[k] >= q.pq[j] {
      break
    }

    q.exchange(k, j)
    k = j
  }
}

func (q *SortingTreePQ[T]) exchange(a, b int) {
  q.pq[a], q.pq[b] = q.pq[b], q.pq[a]
}