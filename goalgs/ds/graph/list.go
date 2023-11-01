package graph

import (
  ll "github.com/bd878/goalgs/ds/linkedlist"
)

type AdjacencyList struct {
  elems []*ll.SentinelLLNode[int]
}

func NewList(size int) *AdjacencyList {
  elems := make([]*ll.SentinelLLNode[int], size)
  for i, _ := range elems {
    elems[i] = ll.InitSentinelLL[int]()
  }
  return &AdjacencyList{elems}
}

func (l *AdjacencyList) Insert(pair []int) {
  if !l.Has(pair) {
    i, j := pair[0], pair[1]

    l.elems[i].Insert(ll.NewSentinelNode[int](j))
    l.elems[j].Insert(ll.NewSentinelNode[int](i))
  }
}

func (l *AdjacencyList) Has(pair []int) bool {
  i, j := pair[0], pair[1]

  found := false
  l.elems[i].Traverse(func (n *ll.SentinelLLNode[int]) {
    if n.Item() == j {
      found = true
    }
  })

  return found
}
