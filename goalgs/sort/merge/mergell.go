package merge

import (
  "golang.org/x/exp/constraints"
  ds "github.com/bd878/goalgs/ds/linkedlist"
)

func MergeLL[T constraints.Ordered](a *ds.DumpHeadNode[T], b *ds.DumpHeadNode[T]) *ds.DumpHeadNode[T] {
  head := ds.InitDumpHeadNode[T]()
  c := head

  for ; a != nil && b != nil; {
    if a.Item() < b.Item() {
      c = c.SetNext(a)
      a = a.Next()
    } else {
      c = c.SetNext(b)
      b = b.Next()
    }
  }

  if a == nil {
    c.SetNext(b)
  } else {
    c.SetNext(a)
  }

  return head.Next()
}