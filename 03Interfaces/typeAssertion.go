package main

import (
	"fmt"
	"log"
)

type shape1 interface {
	area() float64
}

type circle1 struct {
	radius float64
}

// Implementing the area method for circle1
func (c circle1) area() float64 {
	return 3.14 * c.radius * c.radius
}

func valrus() {
	// Assigning a circle1 instance to the shape1 interface
	var s shape1 = circle1{radius: 0}
	// Type assertion to check if s is of type circle1
	c, ok := s.(circle1)
	if !ok {
		log.Fatal("s is not a circle1")
	}

	// Access the radius field
	radius := c.radius

	// Print the results
	fmt.Printf("Radius of circle: %v\n", radius)
	fmt.Printf("Area of circle: %v\n", c.area())
}

func main() {
	valrus()
}
