package main

import "fmt"

// binarySearch реализует бинарный поиск в отсортированном массиве.
// Возвращает индекс элемента, если он найден, и -1, если не найден.
// Сложность алгоритма - O(log n)
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {

		mid := low + (high-low)/2

		if arr[mid] == target {
			return mid // Найдено значение
		} else if arr[mid] < target {
			low = mid + 1 // Искать в правой части
		} else {
			high = mid - 1 // Искать в левой части
		}
	}

	return -1 // Значение не найдено
}

func main() {
	numbers := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	target := 12

	index := binarySearch(numbers, target)

	if index != -1 {
		fmt.Printf("Element %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Element %d not found in the array\n", target)
	}
}
