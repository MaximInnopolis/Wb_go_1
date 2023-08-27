package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	dataMap := make(map[string]int)
	var mu sync.Mutex // Определение мьютекса для защиты доступа к map

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < 20; j++ {
				key := fmt.Sprintf("key%d", j)
				mu.Lock() // Блокировка мьютекса перед записью в map
				dataMap[key] = rand.Intn(100)
				mu.Unlock() // Разблокировка мьютекса после записи
			}
		}(i)
	}

	wg.Wait() // Ожидание завершения всех горутин

	fmt.Println(dataMap) // Вывод содержимого мапы
}
