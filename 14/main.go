package main

import "fmt"

func main() {
	var data interface{} // Переменная типа interface{}

	// Присваиваем значение разных типов
	//data = 42 // int
	data = "Hello" // string
	// data = true       // bool
	// data = make(chan int) // channel

	// Определяем тип значения
	switch value := data.(type) {
	case int:
		fmt.Printf("Переменная содержит целое число: %d\n", value)
	case string:
		fmt.Printf("Переменная содержит строку: %s\n", value)
	case bool:
		fmt.Printf("Переменная содержит булево значение: %t\n", value)
	case chan int:
		fmt.Printf("Переменная содержит канал целых чисел\n")
	default:
		fmt.Println("Тип значения неизвестен")
	}
}
