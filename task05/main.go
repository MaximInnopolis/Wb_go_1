package main

import (
	"fmt"
	"time"
)

func getSeconds() int {

	fmt.Print("Enter seconds to end the program: ")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}

	return n
}

func sender(ch chan<- int, count int) {
	for i := 1; i <= count; i++ {
		ch <- i
		time.Sleep(time.Second) // Имитация задержки отправки
	}
	close(ch)
}

func main() {
	num := getSeconds() // Определение количества секунд работы программы

	dataChan := make(chan int) // Канал для передачи данных

	go sender(dataChan, 10) // Запуск горутины для отправки данных

	timeout := time.After(time.Duration(num) * time.Second)

	for {
		select {
		case data, ok := <-dataChan:
			if !ok {
				fmt.Println("Channel is closed. Exiting...")
				return
			}
			fmt.Printf("Receiving data...\n")
			time.Sleep(time.Second) // Используется сон для имитации получения данных
			fmt.Printf("Received data: %d\n", data)
		case <-timeout:
			fmt.Printf("Timeout of %d seconds reached. Exiting...\n", num)
			return
		}
	}
}
