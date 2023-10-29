package main

import "fmt"

func main() {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7}
	if ind, ok := binarySearch(arr1, 4); ok {
		fmt.Println(ind)
	} else {
		fmt.Println("Значение отсутсвует")
	}

}

func binarySearch(list []int, item int) (int, bool) {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + high
		guess := list[mid]
		if guess == item {
			return mid, true
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}
