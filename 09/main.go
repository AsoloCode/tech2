package main

import "fmt"

func main() {
	ch := make(chan int)
	cnRes := make(chan int)
	arr := [5]int{1, 2, 3, 4, 5}

	go func() {
		for _, num := range arr {
			ch <- num
		}
		close(ch) // Закрываем канал после отправки данных
	}()

	go func() {
		for num := range ch {
			cnRes <- num * 2
		}
		close(cnRes) // Закрываем cnRes, когда все данные обработаны
	}()

	// Чтение результатов из канала cnRes
	for res := range cnRes {
		fmt.Println(res)
	}
}
