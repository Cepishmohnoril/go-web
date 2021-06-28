package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	sqr := square{side: 5}
	cir := circle{radius: 10}

	printArea(sqr)
	printArea(cir)
}
