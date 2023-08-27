package main

import (
	"fmt"
)

func main() {
	temperatures := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создание карты для хранения групп температур
	temperatureGroups := make(map[int][]float32)

	// Заполнение карты группами температур
	for _, temp := range temperatures {
		group := int(temp/10.0) * 10                                      // Определение темп группы
		temperatureGroups[group] = append(temperatureGroups[group], temp) // Добавление значения в темп группу
	}

	// Вывод результата
	for temp, group := range temperatureGroups {
		fmt.Printf("%d: %v\n", temp, group)
	}
}
