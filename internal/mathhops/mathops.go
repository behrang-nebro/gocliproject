package mathops

import (
	"fmt"
	"math"
)

func Add(x, y float64) float64 {
	// fmt.Println("this code id runing from ./internal/mathops/")
	return x + y
}
func Sub(x, y float64) float64 {
	// fmt.Println("this code id runing from ./internal/mathops/")
	return x - y
}
func Mul(x, y float64) float64 {
	// fmt.Println("this code id runing from ./internal/mathops/")
	return x * y
}
func Div(x, y float64) float64 {
	// fmt.Println("this code id runing from ./internal/mathops/")
	if y == 0 {
		fmt.Println("cant divide by zero")
		return 0
	}
	return x / y
}
func Sin(x float64) float64 {
	// fmt.Println("this code id runing from ./internal/mathops/")
	return math.Sin((x / 180) * math.Pi)
}
func Cosin(x float64) float64 {
	// fmt.Println("this code id runing from ./internal/mathops/")
	return math.Cos((x / 180) * math.Pi)
}
