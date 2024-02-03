package bst_test

import (
  "testing"
  "math/rand"

  "github.com/bd878/goalgs/search/bst"
  index "github.com/bd878/goalgs/search/index"
)

func TestBST(t *testing.T) {
  tree := bst.NewBinaryST[uint, *index.STItem]()

  items := make([](*index.STItem), rand.Intn(int(index.MAX_KEY)))
  for i := 0; i < len(items); i++ {
    items[i] = index.NewItem()
    items[i].SetKey(uint(i))
    items[i].SetValue(rand.Float32())
    tree.Insert(items[i])
  }

  item := tree.Search(items[0].Key())
  if item.Key() != items[0].Key() {
    t.Errorf("found wrong item")
  }
}