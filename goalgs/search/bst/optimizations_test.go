package bst_test

import (
  "testing"
  // "math/rand"

  "github.com/bd878/goalgs/search/bst"
)

func makeTree(from, to int) *bst.BinaryST[int, *bst.IntItem] {
  tree := bst.NewBinaryST[int, *bst.IntItem]()

  for i := from; i < to; i++ {
    item := &bst.IntItem{}
    item.Rand()
    item.SetKey(i)
    tree.RandomInsert(item)
  }

  return tree
}

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
  treeA := makeTree(0, 5)
  aCount := treeA.Head().N
  treeB := makeTree(6, 10)
  bCount := treeB.Head().N

  treeA.RandomJoin(treeB)
  if treeA.Head().N != aCount + bCount {
    t.Errorf("wrong count items")
  }
}

func TestBSTRandomJoinLR(t *testing.T) {
  commonTree := makeTree(0, 9)
  if commonTree.Head().L == nil {
    commonTree.TopRotateL()
  }
  if commonTree.Head().R == nil {
    commonTree.TopRotateR()
  }

  treeA := bst.NewBinaryST[int, *bst.IntItem]()

  treeA.Init(commonTree.Head().L)
  aCount := treeA.Head().N

  treeB := bst.NewBinaryST[int, *bst.IntItem]()
  treeB.Init(commonTree.Head().R)
  bCount := treeB.Head().N

  result := treeA.RandomJoinLR(treeA.Head(), treeB.Head())

  treeC := bst.NewBinaryST[int, *bst.IntItem]()
  treeC.Init(result)
  if treeC.Head().N != aCount + bCount {
    t.Errorf("wrong count, got: %d, expected: %d", treeC.Head().N, aCount + bCount)
  }
}

func TestBSTRandomRemove(t *testing.T) {
  tree := makeTree(0, 20)
  prevN := tree.Head().N
  key := tree.Head().Item.Key()
  tree.RandomRemove(key)
  if prevN-1 != tree.Head().N {
    t.Errorf("wrong count, got: %d, expected: %d", tree.Head().N, prevN-1)
  }

  prevN = tree.Head().N
  // remove root
  tree.RandomRemove(tree.Head().Item.Key())
  if prevN-1 != tree.Head().N {
    t.Errorf("root removed, wrong count, got: %d, expected: %d", tree.Head().N, prevN-1)
  }
}