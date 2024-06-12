package main

import (
	"fmt"
	"slices"
)

func sortColors(nums []int) {
	index := 0
	length := len(nums)
	currentColor := slices.Min(nums)

	for index < length {
		if nums[index] != currentColor {
			for i := index + 1; i < length; i++ {
				if nums[i] == currentColor {
					temp := nums[i]
					nums[i] = nums[index]
					nums[index] = temp
					break
				}
			}
		}
		index += 1
		if slices.Contains(nums[index:], currentColor) == false {
			currentColor += 1
		}
	}

	fmt.Println(nums)
}

func main() {
	nums := []int{2, 0, 1, 1, 2, 0}
	sortColors(nums)
}
