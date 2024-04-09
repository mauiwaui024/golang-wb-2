package main

import (
	"fmt"
	"time"
)

func sender(ch chan<- int) {
	for i := 0; ; i++ {
		ch <- i                 // Отправляем значение в канал
		time.Sleep(time.Second) // Задержка 1 секунда
	}
}

func receiver(ch <-chan int) {
	for {
		value := <-ch // Читаем значение из канала
		fmt.Println("Received:", value)
	}
}

func main() {
	N := 5 // Время выполнения программы в секундах

	ch := make(chan int) // Создаем канал

	// Запускаем горутину, которая будет отправлять значения в канал
	go sender(ch)

	// Запускаем горутину, которая будет читать значения из канала
	go receiver(ch)

	// Ждем N секунд
	time.Sleep(time.Second * time.Duration(N))
}
