package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	wordsList := strings.Split(sentence, " ")

	for i := 0; i < len(wordsList); i++ {
		rootAppended := false
		for _, root := range dictionary {
			if strings.HasPrefix(wordsList[i], root) {
				if !rootAppended {
					rootAppended = true
					wordsList[i] = root
				} else if rootAppended && len(wordsList[i]) > len(root) {
					wordsList[i] = root
				}
			}
		}
	}

	return strings.Join(wordsList, " ")
}

func main() {
	fmt.Println(replaceWords([]string{"cat", "rat", "bat"}, "the cattle was rattled by the battery"))
}
