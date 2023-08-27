package main

import (
	"fmt"
)

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создание хешмапы для множества
	set := make(map[string]bool)
	for _, item := range sequence {
		set[item] = true
	}

	// Получение элементов множества в виде среза
	uniqueItems := make([]string, 0, len(set))
	for item := range set {
		uniqueItems = append(uniqueItems, item)
	}

	fmt.Println("Unique items:", uniqueItems)
}
