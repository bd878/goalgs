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

/* untested */
func (h *LinearProbingHashtable) Expand() *LinearProbingHashtable {
  initFn := func(maxM int) *LinearProbingHashtable {
    m := maxM*2
    return &LinearProbingHashtable{
      Items: make([]hashing.Item, m),
      N: 0,
      M: m,
      maxM: maxM,
    }
  }

  t := initFn(h.M*2)
  for i := 0; i < h.M/2; i++ {
    if h.Items[i] != nil {
      t.Insert(h.Items[i])
    }
  }

  return t
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

func (h *LinearProbingHashtable) Remove(k int) {
  i := hashing.HashInt(k, h.M)
  for h.Items[i] != nil && k != h.Items[i].Key() {
    i = (i+1) % h.M
  }
  if h.Items[i] == nil {
    return
  }
  h.Items[i] = nil
  h.N -= 1
  /* shift all items one elemenet right
     since a hole stops search */
  for j := i+1; h.Items[j] != nil; j = (j+1) % h.M {
    v2 := h.Items[j]
    h.Items[j] = nil
    h.N -= 1
    h.Insert(v2)
  }
}