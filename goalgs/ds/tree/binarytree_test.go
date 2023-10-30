package binarytree_test

import (
  "testing"
  "fmt"
  "math/rand"

  btree "github.com/bd878/goalgs/ds/tree"
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
  root := btree.NewNode[int](10)
  leaf := btree.NewNode[int](11)
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
  tree := btree.Init[rune]()

  elems := getElems(size)
  for _, v := range elems {
    tree = tree.Insert(btree.NewNode[rune](v))
  }

  if tree.CountTotal() != size {
    t.Error("size != count", size, tree.CountTotal())
  }

  t.Log("=== height:", tree.Height())
}

func TestBuildBinaryTree(t *testing.T) {
  size := 10 // rand.Intn(100)

  tree := btree.Init[rune]()

  elems := getElems(size)
  for _, v := range elems {
    tree = tree.Insert(btree.NewNode[rune](v))
  }

  if tree.CountTotal() != size {
    t.Error("size != count", size, tree.CountTotal())
  }

  if tree.IsEmpty() {
    t.Error("tree is empty")
  }

  tree.Print(printRune)
}

func getElems(size int) []rune {
  result := make([]rune, size)
  for i := 0; i < size; i++ {
    result[i] = int32(33+rand.Intn(93))
  }

  return result
}

func printRune(r rune, h int) {
  if r == 0 {
    fmt.Printf("%" + fmt.Sprint(h+3) + "v\n", "*")
  } else {
    fmt.Printf("%" + fmt.Sprint(h+3) + "q\n", r)
  }
}