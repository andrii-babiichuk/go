package main

import "fmt"

type Notifier interface {
	Notify(string) error
}

// MailNotifier реалізовує Notifier за допомогою семантики значення
type MailNotifier struct{}

func (MailNotifier) Notify(msg string) error {
	fmt.Printf("Notification '%s' sent via mail", msg)
	return nil
}

// SMSNotifier реалізовує Notifier за допомогою семантики покажчика
type SMSNotifier struct{}

func (*SMSNotifier) Notify(msg string) error {
	fmt.Printf("Notification '%s' sent via SMS", msg)
	return nil
}

// Інтергація 2-х типів, що мають однаковий створює невизначеність і компілятор видає помилку.
type AlertService struct {
	MailNotifier
	SMSNotifier
	subscription int
}

// Перевизначення методів інтегрованих типів
//func (a *AlertService) Notify(msg string) error {
//	if a.subscription == 1 {
//		return a.SMSNotifier.Notify(msg)
//	}
//	return a.MailNotifier.Notify(msg)
//}

func main() {
	asertService := AlertService{
		MailNotifier: MailNotifier{},
		SMSNotifier:  SMSNotifier{},
		subscription: 1,
	}

	asertService.MailNotifier.Notify("test")
	// asertService.Notify("test") // ambiguous selector asertService.Notify
}
