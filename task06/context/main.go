package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // Объявление отменяемого контекста с таймаутом 3 секунды
	defer cancel()                                                          // Закрытие канала

	go worker(ctx) // Создание горутины

	<-ctx.Done()
	fmt.Println("Timeout of 3 seconds reached. Exiting...")

}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // Немедленное завершение работы программы по истечении таймаута
			return
		default:
			fmt.Println("Worker is doing something...")
			time.Sleep(time.Second) // Используется сон для имитации работы
			fmt.Println("Worker did something...")
		}
	}
}
