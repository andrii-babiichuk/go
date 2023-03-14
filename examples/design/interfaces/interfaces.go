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

func main() {
	// Створення порожньої змінної notifier типу Notifier.
	var notifier1 Notifier
	var notifier2 Notifier

	// Значення всіх змінних типу інтерфейс за замовчуванням дорівнює nil.
	fmt.Println(notifier1, notifier2)

	// notifier1 = SMSNotifier{}
	notifier1 = &SMSNotifier{}

	notifier2 = MailNotifier{}
	// notifier2 = &MailNotifier{}
}
