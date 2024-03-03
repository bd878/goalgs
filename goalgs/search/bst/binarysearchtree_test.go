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

func TestBSTRotateL(t *testing.T) {
  tree := bst.NewBinaryST[uint, *index.STItem]()

  items := make([](*index.STItem), 5)
  for i := 0; i < len(items); i++ {
    items[i] = index.NewItem()
    items[i].SetKey(uint(i))
    items[i].SetValue(float32(i))
    tree.Insert(items[i])
  }

  if tree.Head().Item.Key() != items[0].Key() {
    t.Errorf("wrong head key")
  }
  err := tree.TopRotateL()
  if err != nil {
    t.Fatal(err)
  }
  if tree.Head().Item.Key() != items[1].Key() {
    t.Errorf("wrong head key after rotL")
  }
}

func TestBSTRotateR(t *testing.T) {
  tree := bst.NewBinaryST[uint, *index.STItem]()

  items := make([](*index.STItem), 5)
  for i, j := 0, len(items); i < len(items); i++ {
    items[i] = index.NewItem()
    items[i].SetKey(uint(j))
    items[i].SetValue(float32(j))
    tree.Insert(items[i])

    j -= 1
  }

  if tree.Head().Item.Key() != items[0].Key() {
    t.Errorf("wrong head key")
  }
  err := tree.TopRotateR()
  if err != nil {
    t.Fatal(err)
  }
  if tree.Head().Item.Key() != items[1].Key() {
    t.Errorf("wrong head key after rotR")
  }
}

func TestBSTRotLCount(t *testing.T) {
  tree := bst.NewBinaryST[uint, *index.STItem]()

  first := index.NewItem()
  first.SetKey(0)
  first.SetValue(0)
  tree.Insert(first)

  second := index.NewItem()
  second.SetKey(1)
  second.SetValue(1)
  tree.Insert(second)

  third := index.NewItem()
  third.SetKey(2)
  third.SetValue(2)
  tree.Insert(third)

  err := tree.TopRotateL()
  if err != nil {
    t.Error(err)
  }
}

func TestInsertRoot(t *testing.T) {
  tree := bst.NewBinaryST[uint, *index.STItem]()

  items := make([](*index.STItem), 5)
  for i := 0; i < len(items); i++ {
    items[i] = index.NewItem()
    items[i].SetKey(uint(i))
    items[i].SetValue(float32(i))
    tree.InsertInRoot(items[i])
  }

  if tree.Head().Item.Key() != items[3].Key() {
    t.Errorf("wrong head key")
  }
}

func TestBSTSelect(t *testing.T) {
  tree := bst.NewBinaryST[uint, *index.STItem]()

  items := make([](*index.STItem), 5)
  for i := 0; i < len(items); i++ {
    items[i] = index.NewItem()
    items[i].SetKey(uint(i))
    items[i].SetValue(float32(i))
    tree.InsertInRoot(items[i])
  }

  searchKey := 4
  v := tree.Select(searchKey)
  if v == nil {
    t.Errorf("got nil")
  }
  if v.Key() != uint(searchKey) {
    t.Errorf("wrong key, got %d, expected %d\n", v.Key(), searchKey)
  }
}