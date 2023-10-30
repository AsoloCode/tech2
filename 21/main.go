package main

import "fmt"

// DataService - интерфейс для отправки XML данных
type DataService interface {
	SendXmlData()
}

// XmlDocument - структура для работы с XML документом
type XmlDocument struct {
}

func (doc XmlDocument) SendXmlData() {
	fmt.Println("Отправка XML документа!")
}

// JsonDocument - структура для работы с JSON документом
type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXml() string {
	// В реальности здесь происходило бы преобразование JSON в XML
	return "<xml></xml>"
}

// JsonDocAdapter - адаптер для преобразования JSON данных в XML и отправки
type JsonDocAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter JsonDocAdapter) SendXmlData() {
	// В адаптере происходит вызов метода преобразования JSON в XML
	xmlData := adapter.jsonDocument.ConvertToXml()
	// Затем отправка XML данных
	fmt.Println("Отправка XML данных:", xmlData)
}

func main() {
	// Создаем экземпляр JsonDocument
	jsonDoc := JsonDocument{}

	// Создаем адаптер, передавая ему JsonDocument
	adapter := JsonDocAdapter{jsonDocument: &jsonDoc}

	// Используем DataService интерфейс для отправки XML данных, не зная о JsonDocument
	dataService := adapter

	// Отправляем данные (JsonDocument будет автоматически преобразован в XML)
	dataService.SendXmlData()
}
