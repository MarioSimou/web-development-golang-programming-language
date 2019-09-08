package main

import (
	"fmt"
	"math"
)

type square struct {
	height float64
}
type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func main() {
	s1 := square{height: 4}
	r1 := circle{radius: 2}

	info(s1)
	info(r1)
}

func (s square) area() float64 {
	return math.Pow(s.height, 2)
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s shape) {
	fmt.Printf("Area: %v\n", math.Round((s.area()*100))/100)
}
