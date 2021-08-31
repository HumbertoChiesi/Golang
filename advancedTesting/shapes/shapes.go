package shapes

import (
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width  float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}
