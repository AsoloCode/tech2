package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Privet kak dela"
	fmt.Println(reverseWords(str))

}
func reverseWords(input string) string {
	words := strings.Fields(input) // Разбиваем строку на слова
	reversedWords := make([]string, len(words))

	// Переворачиваем каждое слово и сохраняем его в новый слайс
	for i, j := len(words)-1, 0; i >= 0; i-- {
		reversedWords[j] = words[i]
		j++
	}

	// Объединяем перевернутые слова в итоговую строку
	return strings.Join(reversedWords, " ")
}
