package graph

type AdjacencyMatrix struct {
  elems [][]int
}

func NewMatrix(size int) *AdjacencyMatrix {
  elems := make([][]int, size)
  for i, _ := range elems {
    elems[i] = make([]int, size)
  }

  return &AdjacencyMatrix{elems}
}

func (m *AdjacencyMatrix) Insert(pair []int) {
  i, j := pair[0], pair[1]

  m.elems[i][j] = 1
  m.elems[j][i] = 1
}

func (m *AdjacencyMatrix) Has(pair []int) bool {
  i, j := pair[0], pair[1]
  return m.elems[i][j] == 1
}