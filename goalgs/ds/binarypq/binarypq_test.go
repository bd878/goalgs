package binaryqueue_test

import (
  "testing"
  "fmt"

  bq "github.com/bd878/goalgs/ds/binarypq"
  ds "github.com/bd878/goalgs/ds/stack"
)

func TestSortingTree(t *testing.T) {
  n1 := bq.NewNode[int](10)
  n2 := bq.NewNode[int](11)

  n := n1.Pair(n2)
  if n != n2 {
    t.Errorf("n != n2\n")
  }

  if n.L != n1 {
    t.Errorf("expected n2 %v be the root\n", n.V)
  }

  if n.R != nil {
    t.Errorf("right subtree not empty\n")
  }

  if n1.L != nil {
    t.Errorf("n1.L not nil\n")
  }

  if n1.R != nil {
    t.Errorf("n1.R not nil\n")
  }

  printTree(n, t)

  size := treeSize(n, t)
  if size != 2 {
    t.Errorf("tree size is %d, not %d\n", size, 2)
  }
}

func printTree(r *bq.Node[int], t *testing.T) {
  t.Helper()

  fmt.Printf("%4d\n", r.V)
  fmt.Printf("%3d", r.L.V)
  fmt.Printf("%3c\n", '\u002A')
}

func treeDepth(n *bq.Node[int], t *testing.T) int {
  t.Helper()

  var levels int

}

func treeSize(n *bq.Node[int], t *testing.T) int {
  t.Helper()

  var size int

  s := &ds.ArrStack[*bq.Node[int]]{}
  s.Push(n)
  size += 1
  var top *bq.Node[int]
  var err error

  for i := 0; i < 5 && !s.IsEmpty(); i++ {
    top, err = s.Pop() // remove root from stack
    if err != nil {
      panic(err)
    }

    if top.R != nil {
      s.Push(top.R)
      size += 1
    }
    if top.L != nil {
      s.Push(top.L)
      size += 1
    }
  }

  return size
}