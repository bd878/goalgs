package doublelinkedlist_test

import (
  "testing"
  "math/rand"
  dll "github.com/bd878/goalgs/ds/doublelinkedlist"
)

func TestDoubleLL(t *testing.T) {
  size := rand.Intn(10)
  perm := rand.Perm(size)

  q := dll.New[int]()
  for i := 0; i < size; i++ {
    q.Insert(dll.NewNode[int](perm[i]))
  }

  i := 0
  q.Traverse(func(n *dll.Node[int]) {
    if n.Item() != perm[i] {
      t.Errorf("=== %dth values not equal: %v != %v\n", i, n.Item(), perm[i])
    }
  })
}
