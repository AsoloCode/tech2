package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	numberChannel := make(chan int) //исходное число
	squareChannel := make(chan int) //квадрат числа
	var wg sync.WaitGroup           //для ожидания завершения всех горутин

	for _, num := range numbers {
		wg.Add(1)

		go func(n int) { //создается горутина для каждого числа
			defer wg.Done()         //уменьшает счетчик синхронизации wg
			square := n * n         //квадрат
			numberChannel <- n      //отправляем в канал
			squareChannel <- square //отправляем в канал
		}(num) //передаем значение n в анонимную функцию
	}

	go func() { //горутина которая будет ждать когда завершатся все гоурутины и закроет их
		wg.Wait()
		close(numberChannel)
		close(squareChannel)
	}()

	for num := range numberChannel { //бесконечно читает числа из канала а затем их квадраты
		square := <-squareChannel
		fmt.Printf("Квадрат числа %d: %d\n", num, square)
	}
}
