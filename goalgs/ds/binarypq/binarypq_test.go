package binaryqueue_test

import (
  "testing"
  "math/rand"

  btree "github.com/bd878/goalgs/ds/tree"
  bq "github.com/bd878/goalgs/ds/binarypq"
)

func TestSortingTree(t *testing.T) {
  n1 := btree.NewNode[int](10)
  n2 := btree.NewNode[int](11)

  n := bq.Pair(n1, n2)
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
}

func TestGetMax(t *testing.T) {
  q := bq.NewBinaryPQ[rune](1) // autoresize
  perm := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'}

  for _, v := range perm {
    q.Insert(v)
  }

  v := q.GetMax()
  if v != perm[len(perm)-1] {
    t.Error("res != max", v, perm[len(perm)-1])
  }
}

func TestInsert(t *testing.T) {
  size := 5
  q := bq.NewBinaryPQ[int](4)
  perm := rand.Perm(size)

  for _, v := range perm {
    q.Insert(v)
  }

  v := q.GetMax()
  if v != size-1 {
    t.Errorf("max not 11")
  }
}

func TestJoin(t *testing.T) {
  q1 := bq.NewBinaryPQ[int](1)
  q2 := bq.NewBinaryPQ[int](1)

  perm1 := []int{1, 2}
  perm2 := []int{3, 4, 5}

  for _, v := range perm1 {
    q1.Insert(v)
  }
  for _, v := range perm2 {
    q2.Insert(v)
  }

  q1.Join(q2)

  if q1.IsEmpty() {
    t.Error("q1 is empty")
  }
}
