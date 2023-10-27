package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "cat", "tree"}
	var res []string
	hashMap := make(map[string]struct{})

	// Итерируемся по массиву и записываем в хеш таблицу ключом элемент массива
	for _, val := range arr {
		hashMap[val] = struct{}{}
	}

	// Через цикл записываем все ключи в res
	for key, _ := range hashMap {
		res = append(res, key)
	}
	fmt.Println(res)
}

//hashMapRes := make(map[string]struct{})
//
//for _, val := range arr {
//
//	if _, ok := hashMap[val]; ok {
//		if _, ok := hashMapRes[val]; !ok {
//			res = append(res, val)
//			hashMapRes[val] = struct{}{}
//			continue
//		}
//	}
//
//	hashMap[val] = struct{}{}
//}
//fmt.Println(res)
