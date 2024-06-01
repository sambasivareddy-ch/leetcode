package main

import (
	"fmt"
	"math"
)

func scoreOfString(s string) int {
	score, string_length := 0, len(s)

	for i := 0; i < string_length-1; i++ {
		score += int(math.Abs(float64(s[i]) - float64(s[i+1])))
	}

	return score
}

func main() {
	var s string
	fmt.Print("Enter a string:")
	fmt.Scan(&s)

	score := scoreOfString(s)

	fmt.Println("Score: ", score)
}
