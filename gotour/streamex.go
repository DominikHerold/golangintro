package main

import (
	"fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

func main() {
	fmt.Println("Test")
	reader.Validate(MyReader{})
}

func (r MyReader) Read(p []byte) (int, error) {
	for i := 0; i < len(p); i++ {
		p[i] = 'A'
	}

	return len(p), nil
}
