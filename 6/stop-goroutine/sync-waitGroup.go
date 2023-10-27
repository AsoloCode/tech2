package stop_goroutine

import (
	"fmt"
	"sync"
	"time"
)

//Использование sync.WaitGroup:
//можете использовать sync.WaitGroup, чтобы дождаться завершения выполнения горутин.
//В этом случае, горутины должны сами уменьшать счетчик WaitGroup, и главная горутина вызывает Wait() для ожидания.

//var wg sync.WaitGroup
//for i := 0; i < 5; i++ {
//wg.Add(1)
//go func() {
//defer wg.Done()
//
//}()
//}
//									// Для ожидания завершения всех горутин: wg.Wait()

func SyncWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Горутина %d начала работу\n", i)
			// Работа
			time.Sleep(time.Second)
			fmt.Printf("Горутина %d завершила работу\n", i)
		}(i)
	}

	// Ждем, пока все горутины завершат работу
	wg.Wait()

	fmt.Println("Все горутины завершили работу. Программа завершена.")
}
