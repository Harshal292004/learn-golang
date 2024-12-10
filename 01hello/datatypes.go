package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe    bool       = false
	MathInt uint64     = 1<<64 - 1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func mainnn() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
}
