package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func printShape(s shape) {
	var i = 1
	fmt.Printf("Index: %d - Area:  %f - Perimeter: %f \n", i, s.area(), s.perimeter())
}

func main2() {
	var c = circle{
		radius: 90.0,
	}
	printShape(c)

}
