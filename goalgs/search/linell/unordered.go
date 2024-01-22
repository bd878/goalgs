package linell

import (
  "fmt"
  "golang.org/x/exp/constraints"
  ll "github.com/bd878/goalgs/ds/linkedlist"
  "github.com/bd878/goalgs/search/types"
)

/*
 * TODO: Still does not handle duplicate items with keys
 */

type UnorderedST[K constraints.Ordered, I types.Item[K]] struct {
  head ll.LLNode[I]
  n int
  zero I
}

func NewUnorderedST[K constraints.Ordered, I types.Item[K]]() *UnorderedST[K, I] {
  head := ll.InitPtrLL[I]()
  var zero I
  return &UnorderedST[K, I]{head: head, zero: zero, n: 0}
}

func (s *UnorderedST[K, I]) Count() int {
  return s.n
}

func (s *UnorderedST[K, I]) Search(key K) I {
  var res I
  found := false
  s.head.Traverse(func(n ll.LLNode[I]) {
    if found { return; }
    if n.Item().Key() == key {
      res = n.Item()
      found = true
    }
  })
  return res
}

func (s *UnorderedST[K, I]) Insert(x I) {
  s.head.Insert(ll.NewPtrNode[I](x))
  s.n += 1
}

func (s *UnorderedST[K, I]) searchR(x I) ll.LLNode[I] {
  p := s.head
  found := false
  s.head.Traverse(func(n ll.LLNode[I]) {
    if found { return; }
    if n.Item().Key() == x.Key() { found = true; return; }
    p = n
  })
  return p
}

func (s *UnorderedST[K, I]) Remove(x I) {
  prev := s.searchR(x)
  if prev != nil {
    if prev.Item().Key() == x.Key() {
      // then, it is a head
      s.head = s.head.Next()
    } else {
      prev.DeleteNext()
    }
    s.n -= 1
  }
}

func (s *UnorderedST[K, I]) Select(n int) I {
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

func (s *UnorderedST[K, I]) Swap(l, r ll.LLNode[I]) {
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

func (s *UnorderedST[K, I]) Less(l, r ll.LLNode[I]) bool {
  return l.Item().Key() < r.Item().Key()
}

func (s *UnorderedST[K, I]) Sort() {
  /* it is unordered */
}

func (s *UnorderedST[K, I]) Print() {
  s.head.Traverse(func(n ll.LLNode[I]) {
    fmt.Printf("%v ", n.Item().Key())
  })
  fmt.Println()
}