package binarytree_test

import (
  "testing"
  "math/rand"

  tree "github.com/bd878/goalgs/ds/tree"
)

func TestBinaryTree(t *testing.T) {
  for screnario, fn := range map[string] func(*testing.T) {
    "CountTotal": TestCountTotal,
    "BuildBinaryTree": TestBuildBinaryTree,
    "TreeHeight": TestTreeHeight,
  } {
    t.Run(screnario, fn)
  }
}

func TestCountTotal(t *testing.T) {
  root := tree.NewNode[int](10)
  leaf := tree.NewNode[int](11)
  root.L = leaf

  if root.CountTotal() != 2 {
    t.Fatal("total nodes != 2")
  }

  if leaf.CountTotal() != 1 {
    t.Fatal("total nodes != 1")
  }

  if root.CountTotalRecursive() != root.CountTotal() {
    t.Fatal("count total recursive != count total")
  }
}

func TestTreeHeight(t *testing.T) {
  // TODO: measure time for input 10e3, 10e4, 10e5 ...etc
  size := rand.Intn(10e5)
  head := tree.Init[rune]()
  root := head

  elems := getElems(size)
  for _, v := range elems {
    head = head.Insert(tree.NewNode[rune](v))
  }

  if root.CountTotal() != size {
    t.Error("size != count", size, root.CountTotal())
  }

  t.Log("=== height:", root.Height())
}

func TestBuildBinaryTree(t *testing.T) {
  size := 10 // rand.Intn(100)

  head := tree.Init[rune]()
  root := head

  elems := getElems(size)
  for _, v := range elems {
    head = head.Insert(tree.NewNode[rune](v))
  }

  if root.CountTotal() != size {
    t.Error("size != count", size, root.CountTotal())
  }

  if root.IsEmpty() {
    t.Error("tree is empty")
  }

  root.Print(tree.PrintRune)
}

func getElems(size int) []rune {
  result := make([]rune, size)
  for i := 0; i < size; i++ {
    result[i] = int32(33+rand.Intn(93))
  }

  return result
}