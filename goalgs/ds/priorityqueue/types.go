package priorityqueue

import (
  "golang.org/x/exp/constraints"
)

type PriorityQueue[T constraints.Ordered] interface {
  Insert(item T)
  GetMax() T
  Empty() bool
}