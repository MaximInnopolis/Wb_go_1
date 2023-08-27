package main

import (
	"fmt"
	"strings"
)

// Проверка, что все символы в строке уникальные
func areAllCharactersUnique(input string) bool {
	seen := make(map[rune]bool)              // Мапа для отслеживания уникальных символов
	lowerCaseInput := strings.ToLower(input) // Преобразование строки в нижний регистр

	// Переборка символов строки
	for _, char := range lowerCaseInput {
		if seen[char] {
			return false // Если символ уже виден, строка не уникальна
		}
		seen[char] = true // Пометка символа как уже встречаемый
	}

	return true // Если все символы уникальны, возвращается true
}

func main() {
	stringsToCheck := []string{"abcd", "abCdefAaf", "aabcd"}

	// Проверяем каждую строку и выводим результат
	for _, s := range stringsToCheck {
		fmt.Printf("%s — %t\n", s, areAllCharactersUnique(s))
	}
}
