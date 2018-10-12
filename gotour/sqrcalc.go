package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test")

	calc := Sqrt(9)
	fmt.Println(calc)
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * x)
	}

	return z
}
