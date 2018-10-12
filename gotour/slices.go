package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

func main() {
	fmt.Println("Test")
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)

	for i := 0; i < dy; i++ {
		a[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			a[i][j] = uint8(i * j)
		}
	}

	return a
}
