package main

import (
	"flag"
	"fmt"

	mathops "github.com/behrang-nebro/gocliproject/internal/mathhops"
)

func main() {
	// fmt.Println("hello this is calc main")
	// result := mathops.Add(10, 3.6)
	// fmt.Println("the add operation:", result)
	// result = mathops.Sub(10, 4.75)
	// fmt.Println("the sub operation:", result)
	// result = mathops.Mul(3.65, 4)
	// fmt.Println("the mul operation:", result)
	// result = mathops.Div(23.21, 5.6598)
	// fmt.Println("the div operation:", result)
	// result = mathops.Sin(30)
	// fmt.Println("the sin operation:", result)
	// result = mathops.Cosin(60)
	// fmt.Println("the cosin operation:", result)
	op := flag.String("op", "", "the operations: add, sub, mul, div, sin, cos")
	x := flag.Float64("x", 0, "first number")
	y := flag.Float64("y", 0, "second number")

	flag.Parse()

	var result float64

	switch *op {
	case "add":
		result = mathops.Add(*x, *y)
	case "sub":
		result = mathops.Sub(*x, *y)
	case "mul":
		result = mathops.Mul(*x, *y)
	case "div":
		result = mathops.Div(*x, *y)
	case "sin":
		result = mathops.Sin(*x)
	case "cos":
		result = mathops.Cosin(*x)
	default:
		fmt.Println("Unknown operation: ", *op)
	}

	fmt.Println("result:", result)
}
