package main

import (
	"fmt"
)

func setBit(n int64, pos int, bit int) int64 {
	if bit == 1 {
		// Устанавливаем i-й бит в 1 с помощью битового ИЛИ
		return n | (1 << pos)
	} else if bit == 0 {
		// Устанавливаем i-й бит в 0 с помощью инвертирования и битового И
		return n &^ (1 << pos)
	} else {
		// Вернуть исходное число, если значение бита не равно 0 или 1
		return n
	}
}

func main() {
	var num int64 = 69 // Некоторое исходное число
	fmt.Printf("Исходное число: %b\n", num)

	// Установить 3-й бит в 1
	num = setBit(num, 3, 1)
	fmt.Printf("Установлен 3-й бит в 1: %b\n", num)

	// Установить 2-й бит в 0
	num = setBit(num, 2, 0)
	fmt.Printf("Установлен 2-й бит в 0: %b\n")
}
