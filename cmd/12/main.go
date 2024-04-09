package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// опять мутим мапу
	set := make(map[string]bool)

	// Добавление элементов последовательности в множество
	for _, str := range sequence {
		set[str] = true
	}

	// Вывод множества
	fmt.Println("Собственное множество:", set)

}
