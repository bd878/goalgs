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
    pq: make([]T, maxn),
  }
}

func (q *SortingTreePQ[T]) SetItems(pq []T) {
  q.pq = pq
}

func (q *SortingTreePQ[T]) Items() []T {
  return q.pq
}

func (q *SortingTreePQ[T]) SetN(n int) {
  q.n = n
}

func (q *SortingTreePQ[T]) N() int {
  return q.n
}

func (q *SortingTreePQ[T]) Empty() bool {
  return q.n == 0
}

func (q *SortingTreePQ[T]) Insert(item T) {
  q.pq[q.n] = item
  q.fixUp(q.n)
  q.n += 1
}

func (q *SortingTreePQ[T]) GetMax() T {
  max := q.pq[0]
  q.exchange(0, q.n-1)
  q.n -= 1
  q.fixDown(0)

  return max
}

func (q *SortingTreePQ[T]) fixUp(k int) {
  for p := parent(k); k > 0 && q.pq[p] < q.pq[k]; p = parent(k) {
    q.exchange(k, p)
    k = p
  }
}

func (q *SortingTreePQ[T]) FixUp(k int) {
  q.fixUp(k)
}

func (q *SortingTreePQ[T]) fixDown(k int) {
  for l, r := children(k); l < q.n; l, r = children(k) {
    j := l
    if r < q.n && q.pq[l] < q.pq[r] {
      j = r
    }

    if q.pq[k] >= q.pq[j] {
      break
    }

    q.exchange(k, j)
    k = j
  }
}

func (q *SortingTreePQ[T]) FixDown(k int) {
  q.fixDown(k)
}

func (q *SortingTreePQ[T]) exchange(a, b int) {
  q.pq[a], q.pq[b] = q.pq[b], q.pq[a]
}

func (q *SortingTreePQ[T]) Exchange(a, b int) {
  q.exchange(a, b)
}

func parent(i int) int {
  return (i-1)/2
}

func children(i int) (int, int) {
  return 2*i+1, 2*i+2
}