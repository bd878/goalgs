package separatell

import (
  "github.com/bd878/goalgs/search/hashing"
)

type HashtableLL struct {
  Heads []*HashtableNode
  M int
  N int
  MaxM int
}

type HashtableNode struct {
  next *HashtableNode
  Item hashing.Item
}

func NewHashtableLL(maxM, nodesPerList int) *HashtableLL {
  m := int(maxM / nodesPerList)
  return &HashtableLL{
    M: m,
    Heads: make([]*HashtableNode, m),
  }
}

func (s *HashtableLL) Insert(v hashing.Item) {
  i := hashing.HashInt(v.Key(), s.M)
  s.Heads[i] = &HashtableNode{Item: v, next: s.Heads[i]}
  s.N++
}

func (s *HashtableLL) Search(k int) (hashing.Item, error) {
  bucket := s.Heads[hashing.HashInt(k, s.M)]
  return searchR(k, bucket)
}

func searchR(k int, node *HashtableNode) (hashing.Item, error) {
  if node == nil {
    return nil, hashing.ErrNoItem
  }
  if node.Item.Key() == k {
    return node.Item, nil
  }
  return searchR(k, node.next)
}