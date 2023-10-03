package flavius

import (
  ds "github.com/bd878/goalgs/ds/linkedlist"
)

// last is winner
func Run(n, m int) []int {
  result := make([]int, 0, n)

  head := ds.NewCyclicNode[int](1)
  x := head
  for i := 2; i <= n; i++ {
    x = x.Insert(ds.NewCyclicNode[int](i))
  }

  for ; x != x.Next(); {
    for i := 1; i < m; i++ {
      x = x.Next()
    }
    node := x.DeleteNext()
    result = append(result, node.Item())
  }

  result = append(result, x.Item())

  return result
}