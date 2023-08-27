package main

import (
	"fmt"
)

func reverseString(input string) string {
	// Преобразование строки в последовательность рун (Unicode символов)
	runes := []rune(input)

	// Смена местами первой руны с последней, второй с предпоследней и т.д.
	for i := 0; i < len(runes)/2; i++ {
		j := len(runes) - i - 1
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) // Преобразование рун обратно в строку
}

func main() {
	input := "главрыба"
	fmt.Println("Initial:", input)
	fmt.Println("Reversed:", reverseString(input))
}
