package closestpoints

import (
  "fmt"

  ds "github.com/bd878/goalgs/ds/point"
  ll "github.com/bd878/goalgs/ds/linkedlist"
)

type Grid struct {
  // square g*g
  grid [][]*ll.SentinelLLNode[*ds.Point]
  g int
  d float32
  n int
  count int
}

func NewGrid(n int, d float32) *Grid {
  g := int(1/d)
  grid := make([][]*ll.SentinelLLNode[*ds.Point], g)
  for i, _ := range grid {
    row := make([]*ll.SentinelLLNode[*ds.Point], g)
    grid[i] = row
    for j, _ := range row {
      row[j] = ll.InitSentinelLL[*ds.Point]()
    }
  }

  return &Grid{grid, g, d, n, 0}
}

func (g *Grid) Insert(x, y float32) {
  var xg int = int(float32(g.g) * x)
  var yg int = int(float32(g.g) * y)

  var p *ds.Point = &ds.Point{X: x, Y: y}
  g.grid[xg][yg].Insert(ll.NewSentinelNode[*ds.Point](p))

  for i := max(xg-1, 0); i <= min(xg+1, g.g-1); i++ {
    for j := max(yg-1, 0); j <= min(yg+1, g.g-1); j++ {
      row := g.grid[i][j]

      row.Traverse(func(t *ll.SentinelLLNode[*ds.Point]) {
        if p.Distance(t.Item()) < g.d {
          g.count += 1
        }
      })
    }
  }
}

func (g *Grid) CountPoints() int {
  return g.count
}

func (g *Grid) Print() {
  for i, row := range g.grid {
    for j, list := range row {
      fmt.Printf("[%d, %d]: ", i, j)

      if list.IsEmpty() {
        fmt.Println("-")
      } else {
        list.Traverse(func(p *ll.SentinelLLNode[*ds.Point]) {
          fmt.Printf("(%f, %f), ", p.Item().X, p.Item().Y)
        })
      }

      fmt.Println()
    }
  }
}