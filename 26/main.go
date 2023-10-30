package main

import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	charMap := make(map[rune]struct{})
	lowerCaseStr := strings.ToLower(str)

	for _, char := range lowerCaseStr {
		// Если символ уже встречен, строка не уникальна
		if _, ok := charMap[char]; ok {
			return false
		}
		charMap[char] = struct{}{}
	}
	return true
}

func main() {
	fmt.Println(isUnique("abcd"))      // true
	fmt.Println(isUnique("abCdefAaf")) // false
	fmt.Println(isUnique("aabcd"))     // false
}
