package main

import (
	"fmt"
	"time"
)

// sleep - функция для задержки выполнения программы на заданное количество миллисекунд
func sleep(milliseconds int) {
	<-time.After(time.Millisecond * time.Duration(milliseconds))
}

func main() {
	fmt.Println("Начало выполнения")
	sleep(2000) // Задержка на 2 секунды
	fmt.Println("Прошло 2 секунды")
	sleep(3000) // Задержка на 3 секунды
	fmt.Println("Прошло еще 3 секунды")
}
