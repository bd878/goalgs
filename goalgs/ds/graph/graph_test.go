package graph_test

import (
  "testing"

  graph "github.com/bd878/goalgs/ds/graph"
)

func TestGraph(t *testing.T) {
  for scenario, fn := range map[string]func(t *testing.T) {
    "matrix": testAdjacencyMatrix,
    "list": testAdjacencyList,
  } {
    t.Run(scenario, fn)
  }
}

func testAdjacencyMatrix(t *testing.T) {
  size, pairs := genStarGraph()
  m := graph.NewMatrix(size)

  for _, p := range pairs {
    m.Insert(p)
  }

  if m.Has([]int{0, size-1}) {
    t.Error("0 and last must not be linked")
  }
}

func testAdjacencyList(t *testing.T) {
  size, pairs := genStarGraph()
  l := graph.NewList(size)

  for _, p := range pairs {
    l.Insert(p)
  }

  if l.Has([]int{0, size-1}) {
    t.Error("0 and last must not be linked")
  }
}

// star-like graph with a center in the middle
func genStarGraph() (size int, pairs [][]int) {
  size = 2*2+1
  m := size/2

  pairs = make([][]int, size-1)
  for k, i, j := 0, 0, size-1; i != m; i, j = i+1, j-1 {
    pairs[k] = []int{m, i}
    k += 1
    pairs[k] = []int{m, j}
    k += 1
  }

  return size, pairs
}