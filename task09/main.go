package main

import (
	"fmt"
	"sync"
)

func main() {
	input := []int{1, 2, 3, 4, 5}

	// Каналы для передачи данных между этапами конвейера
	firstStage := make(chan int)
	secondStage := make(chan int)

	// WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Этап 1: Пишем числа в первый канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range input {
			firstStage <- num
		}
		close(firstStage)
	}()

	// Этап 2: Умножаем числа на 2 и передаем во второй канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range firstStage {
			secondStage <- num * 2
		}
		close(secondStage)
	}()

	// Этап 3: Выводим результаты из второго канала в stdout
	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range secondStage {
			fmt.Println(result)
		}
	}()

	wg.Wait() // Ожидание завершения всех горутин
}
