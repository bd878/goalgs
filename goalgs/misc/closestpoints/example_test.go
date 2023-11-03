package closestpoints_test

import (
  "fmt"
  "testing"

  misc "github.com/bd878/goalgs/misc/closestpoints"
)

func ExampleCountPoints() {
  points := misc.CountPoints(50, 8.5)
  fmt.Println(len(points))
  // Output: number of found points
}

func TestGrid(t *testing.T) {
  var size int = 10
  var d float32 = 0.1

  points := make([][]float32, size)
  for i := 0; i < size; i++ {
    points[i] = []float32{float32(i)*d, float32(i)*d}
  }

  grid := misc.NewGrid(size, d)
  for _, v := range points {
    grid.Insert(v[0], v[1])
  }
  grid.Print()

  if grid.CountPoints() != size {
    t.Error("must count all points")
  }
}