package main

import (
  "testing"
)

func TestPointDistance(t *testing.T) {
  p1 := &Point{X: 0, Y: 1}
  p2 := &Point{X: 0, Y: 2}

  dist := p1.Distance(p2)
  if dist != 1 {
    t.Errorf("not accurate distance")
  }
}
