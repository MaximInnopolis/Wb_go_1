package main

import (
	"fmt"
)

func main() {
	set1 := []any{2, "two", "four", 4, 6, 8, 10}
	set2 := []any{4, "four", "eight", 8, 12, 16}

	// Создание хешмапы для первого множества
	set1Map := make(map[any]bool)
	for _, value := range set1 {
		set1Map[value] = true
	}
	fmt.Println(set1Map)

	// Создание хешмапы для результата пересечения
	intersectionMap := make(map[any]bool)

	// Заполнение хешмапы результатом пересечения
	for _, value := range set2 {
		if set1Map[value] {
			intersectionMap[value] = true
		}
	}

	// Получение пересечение в виде слайса
	intersection := make([]any, 0, len(intersectionMap))
	for value := range intersectionMap {
		intersection = append(intersection, value)
	}

	fmt.Println("Intersection:", intersection)
}
