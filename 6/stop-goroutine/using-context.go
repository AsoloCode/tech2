package stop_goroutine

import (
	"context"
	"fmt"
	"time"
)

//Использование контекста:
//можно использовать пакет context для управления временем жизни горутин.
//Контексты позволяют передавать сигналы о завершении и устанавливать таймауты.

//ctx, cancel := context.WithCancel(context.Background())
//go func() {
//
//	select {
//	case <-ctx.Done():
//		return 												// Завершаем выполнение горутины при получении сигнала
//	default:
//															// Выполняем работу
//	}
//}()
//															// Для остановки горутины вызываем cancel()

func Context() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершает работу.")
				return
			default:
				fmt.Println("Горутина выполняет работу.")
				// Здесь работа
				time.Sleep(time.Second)
			}
		}
	}()

	// Подождем некоторое время
	time.Sleep(5 * time.Second)

	// Вызываем cancel(), чтобы отправить сигнал остановки в горутину
	cancel()

	// Подождем немного, чтобы убедиться, что горутина завершила свою работу
	time.Sleep(time.Second)

	fmt.Println("Программа завершена.")
}
