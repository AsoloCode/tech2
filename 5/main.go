package main

import (
	"fmt"
	"time"
)

func main() {
	cnSender := make(chan string)
	duration := 3 * time.Second
	go func() {
		timer := time.NewTimer(duration)
		defer close(cnSender)
		for {
			select {
			case <-timer.C:
				fmt.Println("Программа завершена.")
				return
			default:
				cnSender <- "Прошло пол секунды"
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	for msg := range cnSender {
		fmt.Println(msg)
	}
}
