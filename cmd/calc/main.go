package main

import (
	"fmt"

	mathops "github.com/behrang-nebro/calculator/internal"
)

func main() {
	fmt.Println("hello this is calc main")
	result := mathops.Add(2, 3)
	fmt.Println("the add operation:", result)
	result = mathops.Sub(10, 4)
	fmt.Println("the sub operation:", result)
}
