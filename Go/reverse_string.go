package main

import "fmt"

func reverseString(s []byte) {
	n := len(s)

	for i := 0; i < n/2; i++ {
		temp := s[i]
		s[i] = s[n-i-1]
		s[n-i-1] = temp
	}
}

func main() {
	s := []byte{'H', 'E', 'L', 'L', 'O'}

	reverseString(s)

	fmt.Println(s)
}
