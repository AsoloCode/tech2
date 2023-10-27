package main

import "fmt"

func main() {
	a, b := 5, 6

	a = a + b
	b = a - b
	a = a - b

	fmt.Println(a, b)
}

func swap(a, b int) (int, int) {
	return b, a
}
