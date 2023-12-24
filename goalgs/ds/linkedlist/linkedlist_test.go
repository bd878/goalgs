package linkedlist_test

import (
  "testing"
  "math/rand"

  ds "github.com/bd878/goalgs/ds/linkedlist"
)

func TestLinkedLists(t *testing.T) {
  for scenario, fn := range map[string]func(*testing.T){
    "cyclic linked list": TestCyclicLinkedList,
    "dump head node linked list": TestDumpHeadNodeLinkedList,
    "ptr node": TestPtrNode,
  } {
    t.Run(scenario, fn)
  }
}

func TestPtrNode(t *testing.T) {
  perm := rand.Perm(10)

  list := ds.InitPtrLL[int]()
  head := list
  for _, v := range perm {
    list, _ = list.Insert(ds.NewPtrNode[int](v)).(*ds.PtrNode[int])
  }

  if head.IsEmpty() {
    t.Errorf("list is empty")
  }

  i := 0
  head.Traverse(func(tn ds.LLNode[int]) {
    var n, _ = tn.(*ds.PtrNode[int])
    if n.Item() != perm[i] {
      t.Error("item != perm[i]", n.Item(), perm[i])
    }
    i += 1
  })
}

func TestCyclicLinkedList(t *testing.T) {
  val := rand.Intn(100)

  head := ds.NewCyclicNode[int](val)
  if head.Item() != val {
    t.Errorf("head returned wrong value %v\n", head.Item())
  }

  perm := rand.Perm(rand.Intn(100))
  head = ds.NewCyclicNode[int](perm[0])
  next := head
  for i := 1; i < len(perm); i++ {
    next = next.Insert(ds.NewCyclicNode[int](perm[i]))
  }

  i := 0
  head.Traverse(func(v *ds.CyclicNode[int]) {
    if v.Item() != perm[i] {
      t.Errorf("node #%d have wrong value %v\n", i, v.Item())
    }
    i += 1
  })
}

func TestDumpHeadNodeLinkedList(t *testing.T) {
  head := ds.InitDumpHeadNode[int]()

  if head.Next() != nil {
    t.Error("dump head not next is not nil")
  }

  perm := rand.Perm(rand.Intn(100))
  next := head
  for i := 0; i < len(perm); i++ {
    next.Insert(ds.NewDumpHeadNode[int](perm[i]))
    next = next.Next()
  }

  i := 0
  head.Traverse(func (v *ds.DumpHeadNode[int]) {
    if v.Item() != perm[i] {
      t.Errorf("node #%d have wrong value %v\n", i, v.Item())
    }
    i += 1
  })
}