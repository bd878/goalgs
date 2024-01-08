package linell

import (
  "golang.org/x/exp/constraints"
  ll "github.com/bd878/goalgs/ds/linkedlist"
  "github.com/bd878/goalgs/search/types"
)

type SearchTable[K constraints.Ordered, I types.Item[K]] struct {
  head ll.LLNode[I]
  n int
  zero I
}

func New[K constraints.Ordered, I types.Item[K]]() *SearchTable[K, I] {
  head := ll.InitPtrLL[I]()
  var zero I
  return &SearchTable[K, I]{head: head, zero: zero, n: 0}
}

func (s *SearchTable[K, I]) Count() int {
  return s.n
}

func (s *SearchTable[K, I]) Search(key K) I {
  var result I
  s.head.Traverse(func(i ll.LLNode[I]) {
    if i.Item().Key() == key {
      result = i.Item()
    }
  })
  return result
}

func (s *SearchTable[K, I]) Insert(x I) {
  s.head.Insert(ll.NewPtrNode[I](x))
}

func (s *SearchTable[K, I]) searchR(x I) ll.LLNode[I] {
  var prev ll.LLNode[I]
  s.head.Traverse(func(i ll.LLNode[I]) {
    if i.Next() != nil && i.Next().Item().Key() == x.Key() {
      prev = i
    }
  })
  return prev
}

func (s *SearchTable[K, I]) Remove(x I) {
  prev := s.searchR(x)
  if prev != nil {
    prev.DeleteNext()
  }
}

func (s *SearchTable[K, I]) Select(n int) I {
  if n > s.n {
    return s.zero
  }
  var nMin ll.LLNode[I]
  for i := 0; i < n; i++ {
    j := i
    s.head.Traverse(func(i ll.LLNode[I]) {
      if nMin == nil {
        nMin = i
      } else if nMin.Item().Key() > i.Item().Key() {
        if j == 0 {
          nMin = i
        } else {
          j -= 1
        }
      }
    })
  }
  if nMin == nil {
    return s.zero
  }
  return nMin.Item()
}

func (s *SearchTable[K, I]) Swap(l, r ll.LLNode[I]) {
  prevL := s.searchR(l.Item())
  prevR := s.searchR(r.Item())
  if prevL == r {
    prevL.DeleteNext()
    prevR.Insert(l)
    return
  }
  if prevR == l {
    prevR.DeleteNext()
    prevL.Insert(r)
    return
  }

  if prevL != nil {
    prevL.DeleteNext()
    prevL.Insert(r)
  }
  if prevR != nil {
    prevR.Insert(l)
  }
}

func (s *SearchTable[K, I]) Less(l, r ll.LLNode[I]) bool {
  return l.Item().Key() < r.Item().Key()
}

func (s *SearchTable[K, I]) Sort() {
  /* it is unordered */
}
