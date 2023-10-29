package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(remove(arr, 2))

}

// Тут если порядок не важен,
// гораздо более быстрая возможность заменить элемент, который нужно удалить, на элемент в конце фрагмента,
// а затем вернуть n-1 первых элементов
func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Сдвигаем все элементы справа от индекса удаления на единицу влево
// Неэфективно потому что в конечном итоге моэет потребоваться переместить все элементы,а это дорогостоюще
//func remove(slice []int, s int) []int {
//	return append(slice[:s], slice[s+1:]...)
//}
