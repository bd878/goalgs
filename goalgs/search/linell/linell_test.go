package linell_test

import (
  "fmt"
  "testing"
  "math/rand"
  "sort"
  st "github.com/bd878/goalgs/search/linell"
)

type STItem struct {
  k int
  v int
}

func (i *STItem) Key() int {
  return i.k
}

func (i *STItem) Null() bool {
  return false
}

func (i *STItem) Value() int {
  return i.v
}

func (i *STItem) Rand() {
  i.k = rand.Intn(10e4)
  i.v = rand.Intn(10e4)
}

type STItemsCollection []*STItem

func (x STItemsCollection) Len() int {
  return len(x)
}

func (x STItemsCollection) Less(i, j int) bool {
  return x[i].Key() < x[j].Key()
}

func (x STItemsCollection) Swap(i, j int) {
  x[i], x[j] = x[j], x[i]
}

func (x STItemsCollection) Print() {
  print(x)
}

func print(x []*STItem) {
  for _, v := range x {
    fmt.Printf("%v ", v.Key())
  }
  fmt.Println()
}

func TestLLSearchTable(t *testing.T) {
  for scenario, f := range map[string]func(t *testing.T){
    "test unordered st": testUnorderedST,
    "test ordered st": testOrderedST,
  } {
    t.Run(scenario, f)
  }
}

func testUnorderedST(t *testing.T) {
  count := 10

  store := st.NewUnorderedST[int, *STItem]()
  if store.Count() != 0 {
    t.Errorf("store.Count != 0")
  }

  items := make([]*STItem, count)
  for i := 0; i < count; i++ {
    v := &STItem{}
    v.Rand()
    items[i] = v
    store.Insert(v)
  }

  if store.Count() != count {
    t.Error("store.Count !=", count, store.Count())
  }

  for _, v := range items {
    if store.Search(v.Key()) != v {
      t.Error("store.Search() returned wrong item for key:", v.Key())
    }
  }

  for _, v := range items {
    store.Remove(v)
    if store.Search(v.Key()) != nil {
      t.Error("removed but still in store:", v.Key())
    }
  }
  if store.Count() != 0 {
    t.Errorf("removed all items, but store.Count() != 0")
  }
}

func testOrderedST(t *testing.T) {
  count := 10

  store := st.NewOrderedST[int, *STItem]()
  if store.Count() != 0 {
    t.Errorf("store.Count() != 0")
  }

  items := make([]*STItem, count)
  for i := 0; i < count; i++ {
    v := &STItem{}
    v.Rand()
    items[i] = v
    store.Insert(v)
  }

  sorted := STItemsCollection(items)
  sort.Sort(sorted)

  if store.Count() != count {
    t.Error("store.Count !=", count, store.Count())
  }

  for i, v := range sorted {
    got := store.Select(i).Key()
    if got != v.Key() {
      t.Error("store.Select(i), i, returned, required:", i, got, v.Key())
    }
  }

  for _, v := range items {
    store.Remove(v)
    if store.Search(v.Key()) != nil {
      t.Error("removed but still in store:", v.Key())
    }
  }
  if store.Count() != 0 {
    t.Errorf("removed all items, but store.Count() != 0")
  }
}