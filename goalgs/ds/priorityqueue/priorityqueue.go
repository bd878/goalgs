package priorityqueue

import (
  "golang.org/x/exp/constraints"
)

type UnorderedArrayPQ[T constraints.Ordered] struct {
  pq []T
  n int
}

func NewUnorderedArrayPQ[T constraints.Ordered](maxn int) PriorityQueue[T] {
  return &UnorderedArrayPQ[T]{
    pq: make([]T, maxn),
  }
}

func (q *UnorderedArrayPQ[T]) Insert(item T) {
  q.pq[q.n] = item
  q.n += 1
}

func (q *UnorderedArrayPQ[T]) GetMax() T {
  var max int
  for i := 0; i < q.n; i++ {
    if q.pq[max] < q.pq[i] {
      max = i
    }
  }
  q.exchange(max, q.n-1)
  res := q.pq[q.n-1]
  q.n -= 1
  return res
}

func (q *UnorderedArrayPQ[T]) Empty() bool {
  return q.n == 0
}

func (q *UnorderedArrayPQ[T]) exchange(a, b int) {
  q.pq[a], q.pq[b] = q.pq[b], q.pq[a]
}