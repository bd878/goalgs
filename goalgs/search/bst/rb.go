package bst

import (
  "golang.org/x/exp/constraints"
  "github.com/bd878/goalgs/search/types"
)

func insertRB[K constraints.Ordered, I types.Item[K]](h **BTreeNode[I], x I, sw bool) {
  if (*h) == nil {
    return
  }

  hv := *h
  if hv.L != nil && hv.R != nil {
    if hv.L.Red && hv.R.Red {
      // descending, converting child 4-link node
      // to 1 parent Red link
      hv.Red = true    // parent is Red
      hv.L.Red = false
      hv.R.Red = false
    }
  }
  if x.Key() < hv.Item.Key() {
    // must insert left
    if hv.L == nil {
      hv.L = &BTreeNode[I]{Item: x, N: 1, Red: true}
      hv.N += 1
    } else {
      insertRB(&(hv.L), x, false)
    }
    if hv.Red && hv.L.Red && sw {
      // rising, converting two contiguous Red nodes
      // on one 4-link node
      if err := rotR(h); err != nil {
        panic(err)
      }
    }

    // climbing up the tree
    if hv.L != nil && hv.L.L != nil {
      if hv.L.Red && hv.L.L.Red {
        // two contiguious left reds
        if err := rotR(h); err != nil {
          panic(err)
        }
        hv.Red = false
        if hv.R != nil {
          hv.R.Red = true
        }
      }
    }
  } else {
    // must insert right
    if hv.R == nil {
      hv.R = &BTreeNode[I]{Item: x, N: 1, Red: true}
      hv.N += 1
    } else {
      insertRB(&(hv.R), x, true)
    }

    // climbing up the tree
    if hv.Red && hv.R.Red && !sw {
      if err := rotL(h); err != nil {
        panic(err)
      }
    }
    if hv.R != nil && hv.R.R != nil {
      if hv.R.Red && hv.R.R.Red {
        // two contiguous right reds
        if err := rotL(h); err != nil {
          panic(err)
        }
        hv.Red = false
        if hv.L != nil {
          hv.L.Red = true
        }
      }
    }
  }
}

func (s *BinaryST[K, I]) InsertRedBlack(x I) {
  if s.head == nil {
    s.head = &BTreeNode[I]{Item: x, N: 1}
    return
  }
  insertRB(&s.head, x, false)
}