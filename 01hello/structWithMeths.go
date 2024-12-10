package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r rect) area() int {
	return r.width * r.height
}

var r = rect{width: 5, height: 9}

func jacg() {
	fmt.Println(r.area())
}
