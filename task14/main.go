package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Инициализация переменных разных типов
	var a int = 42
	var b string = "Hello"
	var c bool = true
	var d = make(chan int)

	// Создание слайса с переменными интерфейсного типа
	anyTypeArr := []any{a, b, c, d}

	// Итерация по элементам слайса
	for _, v := range anyTypeArr {

		fmt.Println("\nValue: ", v)

		// Определение типа переменной с помощью switch
		switch v.(type) {
		case int:
			fmt.Println("[Switch] Its type: int")
		case string:
			fmt.Println("[Switch] Its type: string")
		case bool:
			fmt.Println("[Switch] Its type: bool")
		case chan int:
			fmt.Println("[Switch] Its type: chan int")
		default:
			fmt.Println("[Switch] Undefined type")
		}

		// Определение типа переменной с помощью reflect
		fmt.Println("[Reflect] Its type:", reflect.TypeOf(v))
	}
}
