package main 

import "fmt"

func threeConsecutiveOdds(arr []int) bool {
	hasThreeConsecutiveOdds := false 
	n := len(arr)

	if n <= 2 {
		return false
	}
	
	for i := 0; i < n-2; i++ {
		if arr[i]%2 == 1 && arr[i+1]%2 == 1 && arr[i+2]%2 == 1{
			hasThreeConsecutiveOdds = true 
			break
		}
	}

	return hasThreeConsecutiveOdds
}

func main() {
	arr := []int{1,2,3,5,7}
	fmt.Println(threeConsecutiveOdds(arr))
}