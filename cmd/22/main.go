package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Инициализация переменных a и b с помощью типа big.Int
	a := big.NewInt(2)
	b := big.NewInt(20)

	// Возводим 2 в степень 20
	twoToTwenty := new(big.Int).Exp(a, b, nil)

	// Выводим результат
	fmt.Println("2^20 =", twoToTwenty)

	// Инициализируем переменные для арифметических операций
	c := new(big.Int).Set(twoToTwenty)
	d := new(big.Int).Set(twoToTwenty)

	// Умножение
	multiplyResult := new(big.Int).Mul(c, d)
	fmt.Println("Умножение:", multiplyResult)

	// Деление
	divideResult := new(big.Int).Div(multiplyResult, c)
	fmt.Println("Деление:", divideResult)

	// Сложение
	addResult := new(big.Int).Add(c, d)
	fmt.Println("Сложение:", addResult)

	// Вычитание
	subtractResult := new(big.Int).Sub(c, d)
	fmt.Println("Вычитание:", subtractResult)

	// Максимальное значение
	// maxInt := new(big.Int).SetUint64(^uint64(0))

	// // Минимальное значение
	// minInt := new(big.Int).Neg(maxInt)

	// fmt.Println("Максимальное значение big.Int:", maxInt)
	// fmt.Println("Минимальное значение big.Int:", minInt)
}
