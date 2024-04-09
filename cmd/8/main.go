package main

import (
	"fmt"
)

func setBit(num int64, pos int, bitValue int) int64 {
	if pos < 0 || pos > 63 {
		fmt.Println("Недопустимое значение индекса бита")
		return num
	}
	if bitValue != 0 && bitValue != 1 {
		fmt.Println("Недопустимое значение бита")
		return num
	}
	if bitValue == 1 {
		num |= (1 << pos)
	} else if bitValue == 0 {
		num &= ^(1 << pos)
	} else {
		fmt.Println("Недопустимое значение бита")
	}

	return num
}

func main() {
	var num int64 = 10
	fmt.Printf("Исходное число: %d\n", num)

	// Установка 2-го бита в 1
	num = setBit(num, 2, 1)
	fmt.Printf("Число после установки 2-го бита в 1: %d\n", num)

}
