package main

import (
	"fmt"

	mathops "github.com/behrang-nebro/calculator/internal/mathhops"
)

func main() {
	fmt.Println("hello this is calc main")
	result := mathops.Add(10, 3.6)
	fmt.Println("the add operation:", result)
	result = mathops.Sub(10, 4.75)
	fmt.Println("the sub operation:", result)
	result = mathops.Mul(3.65, 4)
	fmt.Println("the sub operation:", result)
	result = mathops.Div(23.21, 5.6598)
	fmt.Println("the sub operation:", result)
}
