package bst_test

import (
  "testing"

  "github.com/bd878/goalgs/search/bst"
)

func TestRBBST(t *testing.T) {
  tree := bst.NewBinaryST[int, *bst.IntItem]()

  items := make([]*bst.IntItem, 4)
  for i := 0; i < len(items); i++ {
    items[i] = &bst.IntItem{}
    items[i].Rand()
    items[i].SetKey(i)
    tree.InsertRedBlack(items[i])
  }

  if tree.Count() != len(items) {
    t.Errorf("tree.Count() != len(items): %d, %d\n", tree.Count(), len(items))
  }
}