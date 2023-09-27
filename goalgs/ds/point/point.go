package main

import (
  "math"
)

type Point struct {
  X float32
  Y float32
}

func (p *Point) Distance(b *Point) float32 {
  var dx, dy float32 = p.X-b.X, p.Y-b.Y
  return float32(math.Sqrt(float64(dx*dx+dy*dy)))
}

func (p *Point) Polar() (r float32, theta float32) {
  r = float32(math.Sqrt(float64(p.X*p.X+p.Y*p.Y)))
  theta = float32(math.Atan2(float64(p.Y), float64(p.X)))
  return r, theta
}