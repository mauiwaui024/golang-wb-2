package main

import "fmt"

func findIntersections(set1, set2 []int) []int {
	//слайс с результатом
	intersect := make([]int, 0)

	//делаем мапу для первого сета
	set1Map := make(map[int]bool)

	// закидываем в мапу все элементы первого множества
	for _, val := range set1 {
		set1Map[val] = true
	}

	// чекам мапу на наличие элементов из второго сета
	for _, val := range set2 {
		if set1Map[val] {
			intersect = append(intersect, val)
		}
	}

	return intersect
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	result := findIntersections(set1, set2)

	// Вывод результата
	fmt.Println("Пересечения:", result)
}
