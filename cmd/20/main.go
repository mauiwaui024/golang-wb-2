package main

import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	// Разделяем строку на слова
	words := strings.Fields(input)

	// Переверачиваем порядок слов
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Объединить слова в строку
	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println("reversed string:", reversed)
}
