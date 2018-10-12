package main

import (
	"fmt"
)

func main() {
	fmt.Println("Test")
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		tmp := x
		x = y
		y = tmp + y
		return x
	}
}
