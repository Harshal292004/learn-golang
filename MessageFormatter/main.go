package main

import (
	"fmt"
)

func main() {
	formatter := mf.PlainText()
	input := ` hello , go !`
	output := formatter.Format(input)
	fmt.Println(output)
}
