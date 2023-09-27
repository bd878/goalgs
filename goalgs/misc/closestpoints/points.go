package closestpoints

import (
  "math/rand"

  ds "github.com/bd878/goalgs/ds/point"
)

const MAX_N = 10e2

func CountPoints(n int, d float32) []*ds.Point {
  return countPoints(n, d)
}

func countPoints(n int, d float32) []*ds.Point {
  result := make([]*ds.Point, 0)

  points := make([]*ds.Point, n)
  for i := 0; i < n; i++ {
    points[i] = &ds.Point{
      X: float32(rand.Intn(MAX_N)),
      Y: float32(rand.Intn(MAX_N)),
    }
  }

  for i := 0; i < n; i++ {
    for j := i+1; j < n; j++ {
      dist := points[i].Distance(points[j])
      if dist < d {
        result = append(result, points[i], points[j])
      }
    }
  }

  return result
}