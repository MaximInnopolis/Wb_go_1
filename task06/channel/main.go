package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan struct{}) // Объявление канала
	defer close(channel)           // Закрытие канала

	go worker(channel) // Создание горутины

	select {
	case <-time.After(3 * time.Second): // 3-ех секундное ожидание после чего программа завершается
		channel <- struct{}{}
		fmt.Println("Timeout of 3 seconds reached. Exiting...")
	}
}

func worker(channel chan struct{}) {
	for {
		select {
		case <-channel: // Немедленное завершение работы программы по истечении тайм-аута
			return
		default:
			fmt.Println("Worker is doing something...")
			time.Sleep(time.Second) // Используется сон для имитации работы
			fmt.Println("Worker did something...")
		}
	}
}
