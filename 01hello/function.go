package main

import (
	"fmt"
)

// sharing of data type
// x and y have same dType
func add(x, y int) int {
	return x + y
}

/*
func add(x int, y int)int{
	return x+y
} */

// function can return any number of results
func swap(x, y string) (string, string) {
	return y, x
}

func maimn() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println("The addition of 5,4 is", add(5, 4))
}
