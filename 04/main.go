package main

import (
	"fmt"
	"os"        // для работы с операционной системой.
	"os/signal" // для работы с сигналами, в данном случае, мы будем использовать его для перехвата сигнала Ctrl+C.
	"sync"      // для синхронизации горутин.
	"syscall"   // для работы с системными вызовами, в данном случае, для определения сигналов.
)

func main() {
	numWorkers := 3 // количество воркеров

	dataChannel := make(chan int) // канал для передачи данных из главного потока в воркеры.
	done := make(chan struct{})   //канал для сигнализации о завершении работы воркеро
	var wg sync.WaitGroup         //переменная типа sync.WaitGroup, которую мы будем использовать для ожидания завершения всех воркеров.

	// Запускаем воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		//увеличиваем счетчик ожидания
		go worker(i, dataChannel, done, &wg) //новая горутина которая выполняет функцию worker
	}

	// Запускаем горутину для записи данных в канал
	go func() {
		defer close(dataChannel)

		for i := 1; i <= 10; i++ {
			dataChannel <- i
		}
	}()

	// Ожидаем сигнала завершения (Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-c

	// Завершаем воркеров
	close(done)
	wg.Wait()
}

func worker(id int, dataChannel <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Воркер %d начал работу\n", id)

	for {
		select {
		case data, ok := <-dataChannel: //Если данные доступны в dataChannel, воркер считывает данные из канала и выводит сообщение, указывая, что он получил данные.
			if !ok { //проверка ok, которая проверяет, были ли данные успешно получены из dataChannel.
				fmt.Printf("Воркер %d завершил работу\n", id)
				return
			}
			fmt.Printf("Воркер %d получил данные: %d\n", id, data)
		case <-done: //Если сигнал получен из done, воркер выводит сообщение о завершении работы по сигналу и возвращает из функции.
			fmt.Printf("Воркер %d завершил работу по сигналу\n", id)
			return
		}
	}
}
