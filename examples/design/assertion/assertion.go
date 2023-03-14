package main

import "fmt"

type Notifier interface {
	Notify(string) error
}

type Receiver interface {
	Receive() error
}

type ReceiveNotifier interface {
	Notifier
	Receiver
}

type v1 struct{}

func (v1) Notify(msg string) error {
	fmt.Println("calling notify")
	return nil
}

func (v1) Receive() error {
	fmt.Println("calling receive")
	return nil
}

func main() {
	var rn ReceiveNotifier
	var r Receiver

	r = v1{}

	// У випадку, якщо інтерфейс Receiver не описує всі методи необхідні для ReceiveNotifier
	// можна зробити перевірку типу за допомогою "type assertion".
	// Якщо r містить тип v1 ok буде рівне true а rn міститиме значення типу v1
	rn, ok := r.(v1)
	if ok {
		rn.Receive()
		rn.Notify("")
	}

	// Інтерфейс
	rn2, ok := r.(ReceiveNotifier)
	if ok {
		rn2.Receive()
		rn2.Notify("")
	} else {
		fmt.Println("not ReceiveNotifier")
	}

	// Навіть коли відбувається перевірка використовуючи інтерфейс, перевіряється саме тип значення.
	// У випадку, якщо інтерфейс не містить значення, перевірка видасть негативний результат.
	var rn4 ReceiveNotifier
	rn3, ok := rn4.(ReceiveNotifier)
	if ok {
		rn3.Receive()
		rn3.Notify("")
	} else {
		fmt.Println("not ReceiveNotifier")
	}
}
