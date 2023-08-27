package main

import (
	"fmt"
)

// Сложность алгоритма в худшем случае - O(n^2), где n - кол-во элементов в массиве
// (выбранный элемент оказывается наибольшим или наименьшим элементом в подмассиве,
// что приводит к несбалансированным разбиениям)
// В среднем сложность этого алгоритма - O(n log n)

func quicksort(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quicksort(arr, low, pivotIndex-1)  // Рекурсивная сортировка левой части
		quicksort(arr, pivotIndex+1, high) // Рекурсивная сортировка правой части
	}
}

// Функция partition разделяет слайс на две части относительно опорного элемента
func partition(arr []int, low, high int) int {
	pivot := arr[high] // Опорный элемент
	i := low - 1

	// Проходим по подмассиву и перемещаем элементы меньше опорного влево
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Перемещаем опорный элемент на правильную позицию
	return i + 1
}

func main() {
	numbers := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("Unsorted numbers:", numbers)

	quicksort(numbers, 0, len(numbers)-1) // Вызов функции быстрой сортировки

	fmt.Println("Sorted numbers:", numbers)
}
