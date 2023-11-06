package merge

import (
  "golang.org/x/exp/constraints"
  ds "github.com/bd878/goalgs/ds/linkedlist"
  queue "github.com/bd878/goalgs/ds/queue"
)

func Mergesort(nums []int) {
	if len(nums) > 1 {
		mergesort(nums, 0, len(nums)-1)
	}
}

func MergesortUp(nums []int) {
  if len(nums) > 1 {
    mergesortUp(nums, 0, len(nums)-1)
  }
}

func mergesort(nums []int, from, to int) {
	if from < to {
		half := (from + to) / 2
		mergesort(nums, from, half)
		mergesort(nums, half + 1, to)
		Merge(nums, from, half, to)
	}
}

func mergesortUp(nums []int, from, to int) {
  for m := 1; m <= to-from; m = m*2 {
    // merge to adjacent lists
    for i := 0; i <= to-from; i = i+m*2 {
      ll, mm, rr := i, min(i+m-1, to), min(i+m*2-1, to)
      Merge(nums, ll, mm, rr)
    }
  }
}

func MergesortLL[T constraints.Ordered](c *ds.DumpHeadNode[T]) *ds.DumpHeadNode[T] {
  if c == nil || c.Next() == nil {
    return c
  }

  a, b := c, c.Next()
  // divide list on two even parts
  for ; b != nil && b.Next() != nil; {
    c = c.Next()
    b = b.Next().Next() // 2x step
  }
  b = c.Next()
  c.SetNext(nil)

  return MergeLL[T](MergesortLL[T](a), MergesortLL[T](b))
}

func MergesortLLUp[T constraints.Ordered](c *ds.DumpHeadNode[T]) *ds.DumpHeadNode[T] {
  const max = 10e5
  q := queue.New[*ds.DumpHeadNode[T]](max)

  if c == nil || c.Next() == nil {
    return c
  }

  // fill queue with 1-element lists
  u := ds.InitDumpHeadNode[T]()
  for ; c != nil; c = u {
    u = c.Next()
    c.SetNext(nil)
    q.Enqueue(c)
  }

  c = q.Dequeue()
  for ; !q.IsEmpty(); {
    q.Enqueue(c)
    c = MergeLL[T](q.Dequeue(), q.Dequeue())
  }

  return c
}