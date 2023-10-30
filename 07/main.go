package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем карту (map) для хранения данных
	data := make(map[string]int)
	// Создаем мьютекс (mutex) для синхронизации доступа к карте
	var mu sync.Mutex //спользуется для блокировки доступа к данным из нескольких горутин.

	// Количество горутин для параллельной записи
	numGoroutines := 5
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {

		go func(id int) {
			defer wg.Done()

			// Захватываем мьютекс перед записью в карту
			mu.Lock() // Это предотвращает одновременный доступ к карте из нескольких горутин.
			data[fmt.Sprintf("key%d", id)] = id
			// Освобождаем мьютекс после записи
			mu.Unlock()
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим результат
	fmt.Println("Итоговая карта:")
	fmt.Println(data)
}
