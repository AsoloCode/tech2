package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	resChannel := make(chan int)

	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)

		go func(n int) {
			defer wg.Done()
			square := n * n
			resChannel <- square
		}(num)
	}
	go func() {
		wg.Wait()
		close(resChannel)
	}()
	var sum int
	for num := range resChannel {
		sum += num
	}
	fmt.Printf("Сумма квадратов: %d", sum)
}
