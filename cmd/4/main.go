// Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var workersCount int
	fmt.Println("Введите количество воркеров:")
	fmt.Scanln(&workersCount)
	// Канал для передачи данных от главной горутины к воркерам
	dataCh := make(chan int)
	// Канал для ожидания завершения всех воркеров
	var wg sync.WaitGroup

	// Создаем канал для сигналов завершения
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)

	// Запуск воркеров
	for i := 0; i < workersCount; i++ {
		wg.Add(1)
		go worker(i+1, dataCh, &wg) // Передаем i+1 в качестве уникального идентификатора
	}

	// Главная горутина для записи данных в канал
	go func() {
		defer close(dataCh) // Закрываем канал после завершения записи данных
		for i := 1; ; i++ {
			fmt.Println("send data to channel ", i)
			select {
			case dataCh <- i: // Записываем данные в канал
				time.Sleep(500 * time.Millisecond) // Пауза в полсекунды
			case <-stopCh: // При получении сигнала завершения, выходим из цикла
				return
			}
		}
	}()

	// Ожидание завершения всех воркеров
	wg.Wait()
	fmt.Println("Главная горутина завершила работу.")
}

// Функция воркера
func worker(id int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Сигнализируем о завершении работы воркера
	for data := range dataCh {
		fmt.Printf("Горутина %d получила: %d\n", id, data)
	}
}
