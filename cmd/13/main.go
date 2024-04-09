// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	one := 23
	two := 45
	one, two = easySwitch(&one, &two)
	fmt.Println(one, two)
}

// func changePlaces(first, second *int) (int, int) {
// 	*first = *first + *second  //23+45 = 78
// 	*second = *first - *second // 78 - 45 = 23
// 	//second = 23, first = 78
// 	*first = *first - *second // 78 -23 = 45
// 	return *first, *second
// }

func easySwitch(first, second *int) (int, int) {
	return *second, *first
}
