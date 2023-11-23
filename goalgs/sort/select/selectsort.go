package selectsort

import (
  "sort"

  "golang.org/x/exp/constraints"
  ds "github.com/bd878/goalgs/ds/linkedlist"
)

func Selectsort(nums sort.Interface) {
  for i := 0; i < nums.Len(); i++ {
    min := i
    for j := i; j < nums.Len(); j++ {
      if nums.Less(j, min) {
        min = j
      }
    }
    nums.Swap(min, i)
  }
}

func LLSelectsort[T constraints.Ordered](head *ds.DumpHeadNode[T]) *ds.DumpHeadNode[T] {
  out := ds.InitDumpHeadNode[T]()
  for i := 0; i < 20 && head.Next() != nil; i++ {
    prev := findMaxLL[T](head)
    out.Insert(prev.DeleteNext())
  }
  return out
}

func findMaxLL[T constraints.Ordered](head *ds.DumpHeadNode[T]) *ds.DumpHeadNode[T] {
  var prev *ds.DumpHeadNode[T]

  dump := ds.InitDumpHeadNode[T]()
  dump.SetNext(head)

  dump.Traverse(func(n *ds.DumpHeadNode[T]) {
    if prev == nil {
      prev = n
    }

    if n.Next() != nil {
      if n.Next().Item() > prev.Next().Item() {
        prev = n
      }
    }
  })

  return prev
}