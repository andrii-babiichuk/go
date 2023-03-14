package main

import "fmt"

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

// Інтегрувати можна не лише структуру а і інтерфейс.
type AlertService struct {
	Notifier
	subscription int
}

func main() {
	asertService := AlertService{
		Notifier: MailNotifier{},
		// Notifier:  &SMSNotifier{},
		subscription: 1,
	}

	asertService.Notifier.Notify("test")
	asertService.Notify("test")
}
