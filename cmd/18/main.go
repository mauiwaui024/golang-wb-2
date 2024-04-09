package main

import (
	"fmt"
	"sync"
)

// Counter структура счетчика
type Counter struct {
	mu    sync.Mutex // мьютекс для защиты доступа к счетчику
	value int        // значение счетчика
}

// Inc инкрементирует счетчик
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value возвращает текущее значение счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Количество горутин, которые будут инкрементировать счетчик
	numGoroutines := 1000
	wg.Add(numGoroutines)

	// Запускаем горутины для инкрементирования счетчика
	for i := 0; i < numGoroutines; i++ {
		go func() {
			counter.Inc()
			wg.Done()
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
