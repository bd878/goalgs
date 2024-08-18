package separatell

import (
  "errors"
  "github.com/bd878/goalgs/search/hashing"
)

var ErrNoItem = errors.New("no item")

type Item interface {
  Key() int
  Value() int
}

type HashtableLL struct {
  Heads []*HashtableNode
  M int
  N int
  MaxM int
}

type HashtableNode struct {
  next *HashtableNode
  Item Item
}

func NewHashtableLL(maxM, nodesPerList int) *HashtableLL {
  m := int(maxM / nodesPerList)
  return &HashtableLL{
    M: m,
    Heads: make([]*HashtableNode, m),
  }
}

func (s *HashtableLL) Insert(v Item) {
  i := hashing.HashInt(v.Key(), s.M)
  s.Heads[i] = &HashtableNode{Item: v, next: s.Heads[i]}
  s.N++
}

func (s *HashtableLL) Search(k int) (Item, error) {
  bucket := s.Heads[hashing.HashInt(k, s.M)]
  return searchR(k, bucket)
}

func searchR(k int, node *HashtableNode) (Item, error) {
  if node == nil {
    return nil, ErrNoItem
  }
  if node.Item.Key() == k {
    return node.Item, nil
  }
  return searchR(k, node.next)
}