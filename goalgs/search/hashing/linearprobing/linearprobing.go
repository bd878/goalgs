package linearprobing

import (
  "github.com/bd878/goalgs/search/hashing"
)

type LinearProbingHashtable struct {
  Items []hashing.Item
  M int
  N int
  maxM int
}

func NewLinearProbingHashtable(maxM int) *LinearProbingHashtable {
  m := maxM*2 // maxM - hashtable size
  return &LinearProbingHashtable{
    Items: make([]hashing.Item, m),
    N: 0,
    M: m,
    maxM: maxM,
  }
}

func (h *LinearProbingHashtable) Count() int {
  return h.N
}

func (h *LinearProbingHashtable) Search(k int) (hashing.Item, error) {
  i := hashing.HashInt(k, h.M)
  for h.Items[i] != nil {
    if k == h.Items[i].Key() {
      return h.Items[i], nil
    } else {
      i = (i+1) % h.M
    }
  }
  return nil, hashing.ErrNoItem
}

func (h *LinearProbingHashtable) Insert(v hashing.Item) {
  i := hashing.HashInt(v.Key(), h.M)
  for h.Items[i] != nil {
    i = (i+1) % h.M // searching next non-occupied hole
  }
  h.Items[i] = v
  h.N += 1
}