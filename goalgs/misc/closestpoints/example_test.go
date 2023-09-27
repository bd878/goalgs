package closestpoints_test

import (
  "fmt"
  "flag"

  misc "github.com/bd878/goalgs/misc/closestpoints"
)

var (
  n = flag.Int("ncount", 100, "count points")
  d = flag.Float64("distance", 9.9, "distance between closest points")
)

func ExampleCountPoints() {
  flag.Parse()
  points := misc.CountPoints(*n, float32(*d))
  fmt.Println(len(points))
  // Output: number of found points
}