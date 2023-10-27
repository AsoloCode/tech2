package stop_goroutine

import (
	"fmt"
	"runtime"
	"time"
)

//Использование runtime.Goexit():
//Cпецифический метод, который позволяет завершить выполнение текущей горутины.

//go func() {
//	defer runtime.Goexit() 									// Завершаем текущую горутину
//															// Код
//}()

func RuntimeGoexit() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Printf("Горутина %d начала работу\n", i)
			// Здесь может быть ваша полезная работа
			time.Sleep(time.Second)
			fmt.Printf("Горутина %d завершила работу\n", i)
			runtime.Goexit() // Завершаем текущую горутину
		}(i)
	}

	// Подождем немного, чтобы горутины могли завершить работу
	time.Sleep(3 * time.Second)

	fmt.Println("Программа завершена.")
}
