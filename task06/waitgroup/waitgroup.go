package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup // Создание вэйтгруппы
	wg.Add(1)             // Увеличение счетчика вэйтгруппы на 1

	go worker(&wg) // Создание горутины

	wg.Wait()

}

func worker(wg *sync.WaitGroup) {
	fmt.Println("Worker is doing something...")
	time.Sleep(time.Second) // Используется сон для имитации работы
	fmt.Println("Worker did something...")

	time.Sleep(3 * time.Second)
	fmt.Println("Timeout of 3 seconds reached. Exiting...")

	wg.Done() // Счетчик вэйтгруппы уменьшается
}
