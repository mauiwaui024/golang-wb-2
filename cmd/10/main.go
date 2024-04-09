package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Делаем мапу для группировки температору
	temperatureGroups := make(map[int][]float64)

	step := 10.0

	// Группировка температур
	for _, temp := range temperatures {
		group := int(temp/step) * 10
		temperatureGroups[group] = append(temperatureGroups[group], temp)
	}

	// Вывод результатов
	fmt.Println(temperatureGroups)
}
