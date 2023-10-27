package stop_goroutine

import (
	"fmt"
	"time"
)

//Использование канала для сигнала:
//создаем канал и отправляем сигнал горутине, чтобы она завершила выполнение.
//Обычно используется канал chan struct{} для передачи сигнала без данных.

//stopChan := make(chan struct{})
//go func() {
//	select {
//	case <-stopChan:
//		return                              // Завершаем выполнение горутины при получении сигнала
//	default:
//											// Выполняем работу
//	}
//}()
//											// Для остановки горутины отправляем сигнал в канал: close(stopChan)

func ChannelForSignal() {
	stopChan := make(chan struct{}) // Создаем канал для отправки сигнала остановки
	doneChan := make(chan struct{}) // Создаем канал для сигнала завершения горутины

	// Запускаем горутину
	go func() {
		for {
			select {
			case <-stopChan: // Ожидаем сигнал остановки
				fmt.Println("Горутина завершает работу.")
				doneChan <- struct{}{} // Отправляем сигнал о завершении
				return
			default:
				fmt.Println("Горутина выполняет работу.")
				time.Sleep(time.Second)
			}
		}
	}()

	// Подождем некоторое время
	time.Sleep(5 * time.Second)

	// Отправляем сигнал остановки в горутину
	close(stopChan)

	// Ожидаем завершения горутины
	<-doneChan

	fmt.Println("Программа завершена.")
}
