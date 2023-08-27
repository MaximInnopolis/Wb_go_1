package main

import (
	"fmt"
	"math/big"
)

// Библиотека math/big предоставляет поддержку арифметических операций для работы с большими числами,
// которые не помещаются в стандартные числовые типы Go.

func main() {
	// Инициализация чисел как строки, так как значения больше, чем int64 может содержать
	a := "500000000000000000000"
	b := "25000000000000000000"

	bigA := new(big.Int)
	bigA.SetString(a, 10)

	bigB := new(big.Int)
	bigB.SetString(b, 10)

	//Создание промежуточной переменную, используемая для хранения результатов операций
	result := big.NewInt(0)
	fmt.Println("Addition:", result.Add(bigA, bigB))
	fmt.Println("Subtraction:", result.Sub(bigA, bigB))
	fmt.Println("Multiplication:", result.Mul(bigA, bigB))
	fmt.Println("Division:", result.Div(bigA, bigB))

}
