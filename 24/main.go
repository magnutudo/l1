package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPnt(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}
func (p Point) Dist() float64 {
	dist := math.Abs(math.Abs(p.x) - math.Abs(p.y))
	return dist
}
func main() {
	point := NewPnt(2, 10)
	fmt.Println(point.Dist())

}
