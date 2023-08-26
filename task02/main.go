package main

import (
	"fmt"
	"sync"
)

// Waitgroup - механизм синхронизации горутин

func calculateSquare(number int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения функции

	square := number * number
	fmt.Printf("Square of %d is %d\n", number, square)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1) // Увеличиваем счетчик WaitGroup для каждой горутины
		go calculateSquare(num, &wg)
	}

	// Ждем, пока все горутины завершатся
	wg.Wait()
}
