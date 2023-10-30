package main

import (
	"fmt"
)

func main() {
	a := 1048577 // Пример значения больше 2^20
	b := 2097154 // Пример значения больше 2^20

	// Проверка на то, что a и b больше 2^20
	if a > 1048576 && b > 1048576 {
		// Умножение
		product := a * b
		fmt.Printf("Произведение a и b: %d\n", product)

		// Деление
		division := a / b
		fmt.Printf("Результат деления a на b: %d\n", division)

		// Сложение
		sum := a + b
		fmt.Printf("Сумма a и b: %d\n", sum)

		// Вычитание
		subtraction := a - b
		fmt.Printf("Разность a и b: %d\n", subtraction)
	} else {
		fmt.Println("Значения a и b должны быть больше 2^20.")
	}
}
