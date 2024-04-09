package main

import (
	"fmt"
	"strings"
)

// isUnique возвращает true, если все символы в строке уникальные,
// и false в противном случае. Проверка регистронезависимая.
func isUnique(str string) bool {
	// Приводим строку к нижнему регистру
	str = strings.ToLower(str)

	seen := make(map[rune]bool)
	// Проходим по каждому символу в строке
	for _, char := range str {
		// Если символ уже встречался, возвращаем false
		if seen[char] {
			return false
		}
		// Иначе отмечаем символ как встреченный
		seen[char] = true
	}
	return true
}

func main() {
	testStrings := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, str := range testStrings {
		fmt.Printf("String: %s, IsUnique: %v\n", str, isUnique(str))
	}
}
