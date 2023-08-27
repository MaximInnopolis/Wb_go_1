package main

import (
	"fmt"
	"os"
)

func removeAtIndex(slice []any, index int) []any {
	if index < 0 || index >= len(slice) {
		fmt.Println("Invalid index to remove")
		os.Exit(0)
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []any{"u", "w", "x", "y", "z"}

	fmt.Println("Initial slice:", slice)
	fmt.Println("New slice:", removeAtIndex(slice, 0))
}
