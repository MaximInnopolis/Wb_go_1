package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func getWorkerNum() int {
	fmt.Print("Number of workers: ")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}

	return n
}

func worker(ctx context.Context, id int, dataChan <-chan interface{}, exitChan chan<- interface{}) {

	for data := range dataChan {
		fmt.Printf("Worker %d recieving data...\n", id)
		time.Sleep(time.Second) // Используется сон для имитации получения данных
		fmt.Printf("Worker %d received data: %v\n", id, data)

		select {
		// При завершении контекста горутина завершается
		case <-ctx.Done():
			fmt.Printf("Worker %d has been interrupted by OS signal. Exiting...\n", id)
			time.Sleep(500 * time.Millisecond)
			exitChan <- struct{}{}
			return
		default:
		}
	}
}

func main() {
	numWorkers := getWorkerNum() //Выбор количества воркеров

	// Обраюотка нулевого количества воркеров
	if numWorkers == 0 {
		fmt.Println("Number of workers is 0. Exiting...")
		return
	}

	dataChan := make(chan interface{}) // Объявление канала для данных
	exitChan := make(chan interface{}) // Объявление канала выхода

	ctx, cancel := context.WithCancel(context.Background()) // Обьявление отменяемого контекста
	//  Используется для передачи сигнала завершения выполнения всех операций от главного потока к воркерам
	defer cancel()

	// Запуск воркеров
	for i := numWorkers; i > 0; i-- {
		go worker(ctx, i, dataChan, exitChan)
	}

	// Передача данных в канал
	go func() {
		for i := 1; ; i++ {
			dataChan <- i
		}
	}()
	defer close(dataChan)

	// Обработка сигнала завершения (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		select {
		case <-signalChan:
			cancel() // Отправляем сигнал завершения через контекст
		}
		<-signalChan
		os.Exit(1)
	}()

	<-exitChan
}
