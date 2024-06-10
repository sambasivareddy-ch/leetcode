package main

import (
	"fmt"
	"slices"
)

func heightChecker(heights []int) int {
	expected := slices.Clone(heights)
	slices.Sort(expected)

	nOfIncrctPos := 0

	for i := 0; i < len(heights); i++ {
		if heights[i] != expected[i] {
			nOfIncrctPos += 1
		}
	}

	return nOfIncrctPos
}

func main() {
	heights := []int{1, 1, 2, 1, 3, 4, 5}
	fmt.Println(heightChecker(heights))
}
