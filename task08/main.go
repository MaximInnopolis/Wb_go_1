package main

import (
	"fmt"
	"strconv"
)

func setBit(n int64, bitIndex uint, bitValue int) int64 {
	mask := int64(1) << bitIndex
	// Cоздается маска mask путем сдвига бита со значением 1 на позицию bitIndex с помощью оператора сдвига влево (<<).
	// Это означает, что только i-й бит будет установлен в 1, а все остальные будут равны 0.

	if bitValue == 0 {
		return n &^ mask // Сброс i-го бита в числе n
	}
	return n | mask // Установка i-го бита в числе n
}

func main() {
	var num int64 = 43
	var bitIndex uint
	var bitValue int

	fmt.Printf("Variable %d will be modified.\nEnter bit position: ", num)
	_, err := fmt.Scanln(&bitIndex)
	if err != nil || bitIndex < 0 || bitIndex > 63 { // максимальное количество разрядов для int64 - 63
		fmt.Println("Incorrect data. The bit position to be changed must be in the range from 1 to 63")
		return
	}

	fmt.Print("Enter value of bit (0 or 1): ")
	_, err = fmt.Scanln(&bitValue)
	if err != nil || (bitValue != 0 && bitValue != 1) {
		fmt.Println("Incorrect data. Bit value can only be 1 or 0")
		return
	}

	result := setBit(num, bitIndex, bitValue)
	fmt.Printf("After changing bit result: %d\n", result)

	fmt.Println(num, " in binary notation:", strconv.FormatInt(num, 2))
	fmt.Println(result, "in binary notation:", strconv.FormatInt(result, 2))

}
