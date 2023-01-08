package main

// Адаптер — это структурный паттерн проектирования, который позволяет объектам с несовместимыми интерфейсами работать вместе.
//
//	Адаптер позволяет создать объект-прокладку, который будет превращать вызовы приложения в формат, понятный стороннему классу.
type AnalyticalDataService interface {
	SendXmlData()
}
type XmlDocument struct {
}

func (doc XmlDocument) SendXmlData() {
	println("Отправка XML документа")
}

type JsonDocument struct {
}

func (doc JsonDocument) ConvertToXML() string {
	return "<xml></xml>"
}

type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

// мы используем не сам jsonDocument для конвертации, а прослойку - JsonDocumentAdapter
func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.jsonDocument.ConvertToXML()
	print("Отправка XML данных в сервис аналитики")
}
