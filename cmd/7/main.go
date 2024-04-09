package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.Mutex
	m  map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.m[key]
	return val, ok
}

func main() {
	safeMap := NewSafeMap()

	// Конкурентная запись в map
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}
	wg.Wait()

	// Проверка значений
	for i := 0; i < 10; i++ {
		if val, ok := safeMap.Get(fmt.Sprintf("key%d", i)); ok {
			fmt.Printf("Key: key%d, Value: %d\n", i, val)
		} else {
			fmt.Printf("Key: key%d not found\n", i)
		}
	}
}

// В этом примере SafeMap является структурой, содержащей map[string]int и мьютекс для безопасного доступа к этой map. Методы Set и Get могут быть использованы для установки и получения значений из map соответственно. Используя мьютекс, доступ к map защищается от конкурентного доступа нескольких горутин.
