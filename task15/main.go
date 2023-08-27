package main

import (
	"fmt"
	"strings"
)

// Так как глобальные переменные создаются в куче, а локальные в стеке,
// Лучше всего использовать локальные переменные (Работа со стеком быстрее)

// strings.Repeat
//
// Когда вызываетcя strings.Repeat, он создает strings.Builder,
// который представляет собой буфер для построения строки.
// В этом буфере временно хранятся символы до тех пор, пока не будет создана окончательная строка.
//
// Вместо создания новой строки на каждом повторении символа,
// strings.Builder эффективно добавляет символы в свой буфер,
// избегая лишних копирований и аллокаций памяти.
//
// Используя strings.Builder, можно минимизировать количество аллокаций памяти и
// управлять использованиеsм памяти более эффективно.
// Вместо того чтобы создавать временные строки на каждом шаге повторения,
// strings.Builder создает буфер определенного размера и увеличивает его при необходимости.

// createHugeString создает и возвращает строку заданного размера, состоящую из символов "A".
func createHugeString(size int) string {
	return strings.Repeat("A", size) // Используется strings.Repeat для более эффективного создания строк.
}

func someFunc() {
	// Создается буфер для хранения строки в виде среза байтов.
	// Строка преобразовывается в байты для эффективной работы с ней.
	justString := []byte(createHugeString(1 << 10))

	// Ограничиваем длину среза, чтобы получить подстроку из первых 100 байтов.
	justString = justString[:100]

	// Преобразуем байты обратно в строку и выводим результат.
	fmt.Println(string(justString))
}

func main() {
	someFunc()
}