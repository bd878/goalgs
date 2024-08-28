package doublehashing

import (
  "github.com/bd878/goalgs/search/hashing"
)

type DoubleHashingTable struct {
  Items []hashing.Item
  M int
  N int
  maxM int
}

func NewDoubleHashingTable(maxM int) *DoubleHashingTable {
  m := maxM*2 // maxM - hashtable size
  return &DoubleHashingTable{
    Items: make([]hashing.Item, m),
    N: 0,
    M: m,
    maxM: maxM,
  }
}

/* hash value and M must be coprime integers */
func hashtwo(k int) int {
  return (k % 97) + 1;
}

func (h *DoubleHashingTable) Insert(v hashing.Item) {
  if h.N == h.M {
    panic("double hashing table is full")
  }

  i, j := hashing.HashInt(v.Key(), h.M), hashtwo(v.Key())
  for h.Items[i] != nil {
    i = (i + j) % h.M
  }
  h.Items[i] = v
  h.N += 1
}

func (h *DoubleHashingTable) Search(k int) (hashing.Item, error) {
  i, j := hashing.HashInt(k, h.M), hashtwo(k)
  for h.Items[i] != nil {
    if k == h.Items[i].Key() {
      return h.Items[i], nil
    }
    i = (i + j) % h.M
  }
  return nil, hashing.ErrNoItem
}