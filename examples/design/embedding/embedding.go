package main

import "fmt"

type Notifier interface {
	Notify(string) error
}

// MailNotifier реалізовує Notifier за допомогою семантики значення
type MailNotifier struct {
	Name string
}

func (MailNotifier) Notify(msg string) error {
	fmt.Printf("Notification '%s' sent via mail", msg)
	return nil
}

//type AlertService struct {
//	MailNotifier MailNotifier // звичайне поле
//	subscription int
//}

// Інтергація типу MailNotifier в AlertService.
// Інтеграцію варто застосовувати для перевикористання поведінки, а не стану.
type AlertService struct {
	MailNotifier
	subscription int
}

func (c *AlertService) HasAlerts() bool {
	return true
}

func main() {
	asertService := AlertService{
		MailNotifier: MailNotifier{},
		subscription: 1,
	}

	// Доступ до поля
	asertService.Name = "name"
	fmt.Println(asertService.MailNotifier.Name)

	// Доступ до методу
	asertService.Notify("test")
	asertService.MailNotifier.Notify("test")

	// asertService через інтеграцію MailNotifier реалізовує інтерфейс Notifier
	Notify(asertService, "test")
}

func Notify(notifier Notifier, message string) {
	if err := notifier.Notify(message); err != nil {
		fmt.Println("error")
	}
}
