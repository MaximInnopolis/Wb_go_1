package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика.
type Counter struct {
	value int
	mu    sync.Mutex
}

// Increment - инкрементирует значение счетчика на 1.
func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

// GetValue - возвращает текущее значение счетчика.
func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	numWorkers := 10 // Количество конкурентных горутин

	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // Ожидание завершения всех воркеров

	fmt.Println("Final counter value:", counter.GetValue())
}
