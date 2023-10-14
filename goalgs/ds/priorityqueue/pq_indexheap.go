package priorityqueue

import (
  "golang.org/x/exp/constraints"
)

type IndexHeapPQ[T constraints.Ordered] struct {
  v []T // compare indexes through slice values

  pos []int // position to index
  ids map[int]int // index to position
  n int
}

func NewIndexHeapPQ[T constraints.Ordered](maxn int) *IndexHeapPQ[T] {
  return &IndexHeapPQ[T]{
    pos: make([]int, maxn),
    ids: make(map[int]int, maxn),
  }
}

func (q *IndexHeapPQ[T]) SetValues(v []T) {
  q.v = v
}

// i, j - positions
func (q *IndexHeapPQ[T]) exchange(i, j int) {
  q.pos[i], q.pos[j] = q.pos[j], q.pos[i]

  q.ids[q.pos[i]] = i
  q.ids[q.pos[j]] = j
}

// k - position
func (q *IndexHeapPQ[T]) fixUp(k int) {
  for p := parent(k); k > 0 && q.v[q.pos[p]] < q.v[q.pos[k]]; p = parent(k) {
    q.exchange(k, p)
    k = p
  }
}

// k - position
func (q *IndexHeapPQ[T]) fixDown(k int) {
  for l, r := children(k); l < q.n; l, r = children(k) {
    j := l
    if r < q.n && q.v[q.pos[l]] < q.v[q.pos[r]] {
      j = r
    }

    if q.v[q.pos[k]] >= q.v[q.pos[j]] {
      break
    }

    q.exchange(k, j)
    k = j
  }
}

func (q *IndexHeapPQ[T]) Empty() bool {
  return q.n == 0
}

func (q *IndexHeapPQ[T]) Insert(i int) {
  q.pos[q.n] = i
  q.ids[i] = q.n
  q.fixUp(q.n)
  q.n += 1
}

func (q *IndexHeapPQ[T]) GetMaxIndex() int {
  maxi := q.pos[0]
  q.exchange(0, q.n-1)
  q.n -= 1
  q.fixDown(0)

  return maxi
}

// implied that client changed array value at position i
func (q *IndexHeapPQ[T]) Change(i int) {
  q.fixUp(q.ids[i])
  q.fixDown(q.ids[i])
}

