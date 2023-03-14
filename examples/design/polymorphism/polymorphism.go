package main

import (
	"fmt"
)

// Notifier - Типи interface описують поведінку і не мають стану.
// Якщо можливо, рекомендується робити інтерфейси максимально простими, найкраще на 1 метод.
// Якщо інтерфейс описує 1 метод рекомендується назвати його як [імʼя методу] + "er". Notifier/Reader/Writer.
// У випадку, якщо інтерфейс мітить 2 методи, може бути ReadWriter/ReadCloser.
type Notifier interface {
	Notify(string) error
}

// MailNotifier реалізовує Notifier за допомогою семантики покажчика
type MailNotifier struct{}

func (MailNotifier) Notify(msg string) error {
	fmt.Printf("Notification '%s' sent via mail", msg)
	return nil
}

// SMSNotifier реалізовує Notifier за допомогою семантики значення
type SMSNotifier struct{}

func (*SMSNotifier) Notify(msg string) error {
	fmt.Printf("Notification '%s' sent via SMS", msg)
	return nil
}

// Notify очікує першим аргументом тип, що реалізовує інтерфейс Notifier.
// Тобто будь-який тип, що має метод Notify, що приймає строку та повертає помилку.
func Notify(notifier Notifier, message string) {
	// ... бізнес логіка ...
	if err := notifier.Notify(message); err != nil {
		fmt.Println("error")
	}
}

func main() {
	// Створення порожньої змінної notifier типу Notifier.
	var notifier Notifier // == nil

	notifier = &SMSNotifier{}
	// notifier = MailNotifier{}

	Notify(notifier, "hello")

	//notifiers := []Notifier{
	//	&SMSNotifier{},
	//	MailNotifier{},
	//}
	//for _, n := range notifiers {
	//	Notify(n, "test")
	//}
}
