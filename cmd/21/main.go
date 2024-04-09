// Паттерн "Адаптер" позволяет объектам с несовместимыми интерфейсами работать вместе

package main

import "fmt"

// MessageSender интерфейс для отправки сообщений
type MessageSender interface {
	Send(message string)
}

// ServiceA структура для сервиса A, который использует метод Send()
type ServiceA struct{}

// Send отправляет сообщение
func (s *ServiceA) Send(message string) {
	fmt.Println("Сервис A отправляет сообщение:", message)
}

// ServiceB структура для сервиса B, который использует метод Deliver()
type ServiceB struct{}

// Deliver доставляет сообщение
func (s *ServiceB) Deliver(message string) {
	fmt.Println("Сервис B доставляет сообщение:", message)
}

// ServiceBAdapter адаптер для сервиса B, чтобы он соответствовал интерфейсу MessageSender
type ServiceBAdapter struct {
	serviceB *ServiceB
}

// Send переадресует вызов к методу Deliver() для отправки сообщения
func (a *ServiceBAdapter) Send(message string) {
	a.serviceB.Deliver(message)
}

func main() {
	serviceA := &ServiceA{}
	serviceB := &ServiceB{}
	serviceBAdapter := &ServiceBAdapter{serviceB: serviceB}

	sendMessage(serviceA, "Привет, сервис A!")
	sendMessage(serviceBAdapter, "Привет, сервис B!")
}

// sendMessage принимает интерфейс MessageSender и отправляет сообщение через него
func sendMessage(sender MessageSender, message string) {
	sender.Send(message)
}
