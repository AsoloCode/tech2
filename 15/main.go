package main

/*
Есть потенциальная проблема, связанная с возможным паническим состоянием из-за индексации строки v вне её длины.
Если v короче 100 символов, код вызовет панику.
Чтобы избежать этой проблемы, вам следует проверить, что длина v больше или равна 100 символам перед выполнением индексации

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		// Если v слишком короткая, обрабатываем эту ситуацию, например,назначаем пустую строку.
		justString = ""
	}
}

func main() {
	someFunc()
}
*/

//Пример с задачи
/*
var justString string
func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func main() {
	someFunc()
}

var justString string

*/
