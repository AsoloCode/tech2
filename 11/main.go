package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{3, 4, 4, 7, 8, 5, 6, 6, 8, 5}

	fmt.Println(cross(arr1, arr2))

}

func cross(a1, a2 []int) []int {
	hashMap := make(map[int]bool)
	var res []int

	//Заполняем хеш таблицу где ключ ровняется элементом первого массива, а значение bool
	for _, num := range a1 {
		hashMap[num] = true
	}

	//Бежим по второму массиву и проверяем есть ли у нас элемент второго массива в хеш таблице
	for _, num := range a2 {

		//Если нет значения в хеш таблице, то возвращаемый результат будет всегда false
		if hashMap[num] {
			res = append(res, num)
			hashMap[num] = false
		}
	}

	return res
}
