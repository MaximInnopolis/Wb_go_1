package main

import (
	"fmt"
	"sync"
)

func calculateSquareSum(number int, wg *sync.WaitGroup, resultChan chan<- int) {
	defer wg.Done() // Уменьшаем счетчик WaitGroup после завершения функции

	square := number * number
	resultChan <- square
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	resultChan := make(chan int, len(numbers))

	for _, num := range numbers {
		wg.Add(1) // Увеличиваем счетчик WaitGroup для каждой горутины
		go calculateSquareSum(num, &wg, resultChan)
	}

	// Закрываем канал после завершения всех горутин
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var totalSum int
	for square := range resultChan {
		totalSum += square
	}

	fmt.Printf("Sum of squares: %d\n", totalSum)
}
