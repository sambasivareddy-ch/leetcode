package main

import "fmt"

func appendCharacters(s string, t string) int {
	sIdx, tIdx, sLen, tLen := 0, 0, len(s), len(t)

	for sIdx < sLen && tIdx < tLen {
		if s[sIdx] == t[tIdx] {
			tIdx += 1
		}
		sIdx += 1
	}

	return tLen - tIdx
}

func main() {
	s := "hei"
	t := "hello"

	fmt.Println(appendCharacters(s, t))
}
