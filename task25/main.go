package main

import (
	"fmt"
	"time"
)

// sleep имитирует задержку выполнения программы на указанное количество миллисекунд.
func sleep(seconds time.Duration) {
	<-time.After(seconds * time.Second)
}

func main() {
	fmt.Println("Start sleeping...")

	sleep(3)

	fmt.Println("Woke up")
}
