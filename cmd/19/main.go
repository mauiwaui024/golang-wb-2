package main

import (
	"fmt"
)

// reverseString переворачивает строку, учитывая символы Unicode
func reverseString(input string) string {
	// Переводим строку в массив рун (unicode символов)
	runes := []rune(input)
	length := len(runes)

	// Проходим через половину массива и меняем местами символы
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	return string(runes)
}

func main() {
	input := "главрыба"
	reversed := reverseString(input)
	fmt.Println("Перевернутая строка:", reversed)
}
