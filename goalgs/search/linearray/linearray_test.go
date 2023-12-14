package linearray_test

import (
  "testing"

  "github.com/bd878/goalgs/search/index"
  "github.com/bd878/goalgs/search/linearray"
)

func TestLinearray(t *testing.T) {
  a := linearray.New[uint, *index.STItem]()
  items := make([](*index.STItem), 10)
  for i := 0; i < len(items); i++ {
    items[i] = index.NewItem()
    items[i].Rand()
    a.Insert(items[i])
  }
  if a.Count() != len(items) {
    t.Error("wrong amount of items")
  }
  for i := 0; i < len(items); i++ {
    v := a.Search(items[i].Key())
    if v != items[i] {
      t.Error("search() returned wrong item")
    }
  }
  for i := 0; i < len(items); i++ {
    a.Remove(items[i])
  }
  if a.Count() != 0 {
    t.Error(a.Count(), "!=", 0)
  }
}