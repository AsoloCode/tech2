package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало")
	Sleep(3) // Приостановить выполнение на 3 секунды
	fmt.Println("Прошло 3 секунды")
}
