package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func main() {
	fmt.Println("Test")
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	splitted := strings.Fields(s)
	result := make(map[string]int)
	for i := 0; i < len(splitted); i++ {
		_, ok := result[splitted[i]]
		if ok {
			result[splitted[i]] = result[splitted[i]] + 1
		} else {
			result[splitted[i]] = 1
		}
	}
	return result
}
