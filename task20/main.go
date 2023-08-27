package main

import (
	"fmt"
	"strings"
)

func reverse(str string) string {

	// Преобразование строки в слайс, разделенный пробелами.
	strSlice := strings.Split(str, " ")

	// Смена местами первого слова с последним, второго с предпоследним и т.д.
	for i := 0; i < len(strSlice)/2; i++ {
		j := len(strSlice) - i - 1
		strSlice[i], strSlice[j] = strSlice[j], strSlice[i]
	}

	// Объединение перевернутых слов обратно в строку, используя пробел в качестве разделителя.
	return strings.Join(strSlice, " ")
}

func main() {
	input := "Never gonna give you up"

	fmt.Println("Initial:", input)
	fmt.Println("Reversed:", reverse(input))

}
