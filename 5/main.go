package main

import (
	"fmt"
	"time"
)

func main() {
	cnSender := make(chan string) // канал для отправки сообщений
	duration := 5 * time.Second   // таймер для завершения работы
	go func() {
		timer := time.NewTimer(duration) // устанавливаем таймер
		defer close(cnSender)            // закрываем канал после того как завершится горутина
		for {
			select {
			case <-timer.C: //если таймер срабатывает то выполняется этот кейс
				fmt.Println("Программа завершена.")
				return // завершаем
			default: //если таймер еще не сработал то будет выполняться этот кейс
				cnSender <- "Прошло пол секунды"
				time.Sleep(time.Millisecond * 500) //программа засыпает на 500 милисек. перед отправкой след. сообщения
			}
		}
	}()

	for msg := range cnSender {
		fmt.Println(msg)
	}
}
