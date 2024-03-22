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

func TestBSTRandomJoin(t *testing.T) {
  makeTree := func(from, to int) *bst.BinaryST[int, *bst.IntItem] {
    tree := bst.NewBinaryST[int, *bst.IntItem]()

    for i := from; i < to; i++ {
      item := &bst.IntItem{}
      item.Rand()
      item.SetKey(i)
      tree.RandomInsert(item)
    }

    return tree
  }

  treeA := makeTree(0, 5)
  aCount := treeA.Head().N
  treeB := makeTree(6, 10)
  bCount := treeB.Head().N

  treeA.RandomJoin(treeB)
  if treeA.Head().N != aCount + bCount {
    t.Errorf("wrong count items")
  }
}