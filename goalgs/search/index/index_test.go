package index_test

import (
  "testing"
  "math/rand"

  index "github.com/bd878/goalgs/search/index"
)

func TestIndexSearch(t *testing.T) {
  st := index.New()

  items := make([](*index.STItem), rand.Intn(int(index.MAX_KEY)))
  for i := 0; i < len(items); i++ {
    items[i] = index.NewItem()
    items[i].SetKey(uint(i))
    items[i].SetValue(rand.Float32())
    st.Insert(items[i])
  }

  if st.Count() != len(items) {
    t.Error("count not equal")
  }

  for _, v := range items {
    item := st.Search(v.Key())
    if item != v {
      t.Error("find wrong item")
    }
    st.Remove(item)
  }

  if st.Count() != 0 {
    t.Error("count not zero")
  }
}