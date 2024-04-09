package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Создаем слайс переменных типа интерфейс
	var data []interface{}

	// Добавляем разные данные в слайс
	data = append(data, 42)             // int
	data = append(data, "hello")        // string
	data = append(data, true)           // bool
	data = append(data, make(chan int)) // channel

	// Перебираем элементы слайса и определяем их типы
	for _, item := range data {
		switch item.(type) {
		case int:
			fmt.Println("Это переменная типа int:", item)
		case string:
			fmt.Println("Это переменная типа string:", item)
		case bool:
			fmt.Println("Это переменная типа bool:", item)
		case chan int:
			fmt.Println("Это переменная типа channel:", item)
		default:
			// Если тип неизвестен, выводим его с помощью reflect
			fmt.Println("Неизвестный тип:", reflect.TypeOf(item))
		}
	}
}
