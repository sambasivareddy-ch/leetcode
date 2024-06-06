package main

import (
	"fmt"
	"slices"
)

func isNStraightHand(hand []int, groupSize int) bool {
	handLength := len(hand)

	counter := make(map[int]int)
	unique := make([]int, 0)

	if handLength%groupSize != 0 {
		return false
	}

	for _, j := range hand {
		if _, found := counter[j]; found == true {
			counter[j] += 1
		} else {
			counter[j] = 1
			unique = append(unique, j)
		}
	}

	slices.Sort(unique)

	for _, i := range unique {
		if counter[i] > 0 {
			count := counter[i]
			for j := 0; j < groupSize; j++ {
				if counter[i+j] < count {
					return false
				}
				counter[i+j] -= count
			}
		}
	}

	return true
}

func main() {
	hand := []int{8, 10, 12}
	groupSize := 3
	fmt.Println(isNStraightHand(hand, groupSize))
}
