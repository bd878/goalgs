package bst_test

import (
  "testing"
  // "math/rand"

  "github.com/bd878/goalgs/search/bst"
)

func TestBSTRandomInsert(t *testing.T) {
  tree := bst.NewBinaryST[int, *bst.StringItem]()

  itemsCount := 10
  for i := 0; i < itemsCount; i++ {
    item := &bst.StringItem{} // derived case
    item.Rand()
    item.SetKey(i)
    tree.RandomInsert(item)
  }

  h := tree.Height()
  if h >= itemsCount-1 {
    t.Errorf("tree is derived, must be randomly balanced")
  }
}