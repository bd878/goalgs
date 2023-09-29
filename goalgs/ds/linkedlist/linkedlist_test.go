package linkedlist_test

import (
  "testing"
  "math/rand"
  ds "github.com/bd878/goalgs/ds/linkedlist"
)

func TestLinkedLists(t *testing.T) {
  for scenario, fn := range map[string]func(*testing.T){
    "cyclic linked list": testCyclicLinkedList,
  } {
    t.Run(scenario, fn)
  }
}

func testCyclicLinkedList(t *testing.T) {
  val := rand.Intn(100)

  head := ds.NewCyclicNode[int](val)
  if head.Item() != val {
    t.Errorf("head returned wrong value %v\n", head.Item())
  }

  if !head.IsOnlyOne() {
    t.Error("head is not the only one")
  }

  perm := rand.Perm(rand.Intn(100))
  head = ds.NewCyclicNode[int](perm[0])
  next := head
  for i := 1; i < len(perm); i++ {
    next.Insert(ds.NewCyclicNode[int](perm[i]))
    next = next.Next()
  }

  i := 0
  head.Traverse(func(v *ds.CyclicNode[int]) {
    if v.Item() != perm[i] {
      t.Errorf("node #%d have wrong value %v\n", i, v.Item())
    }
    i += 1
  })
}