package main

import "fmt"

func main() {
	a := []int{1, 4, 3, 2, 6, 7, 8, 5}
	fmt.Println(quickSort(a))
}

func quickSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0] // это опорное число
	var left, right []int

	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}
