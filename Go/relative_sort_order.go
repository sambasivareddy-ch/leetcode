package main

import (
	"fmt"
	"slices"
)

func relativeSortArray(arr1 []int, arr2 []int) []int {
	resultArr := make([]int, 0)
	remainArr := make([]int, 0)

	counter := make(map[int]int)

	for _, j := range arr1 {
		if _, found := counter[j]; found == true {
			counter[j] += 1
		} else {
			counter[j] = 1
		}
	}

	for _, j := range arr2 {
		if count, found := counter[j]; found == true {
			for i := 0; i < count; i++ {
				resultArr = append(resultArr, j)
				delete(counter, j)
			}
		}
	}

	for i, j := range counter {
		for k := 0; k < j; k++ {
			remainArr = append(remainArr, i)
		}
	}

	slices.Sort(remainArr)

	return slices.Concat(resultArr, remainArr)
}

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}

	fmt.Println(relativeSortArray(arr1, arr2))
}
