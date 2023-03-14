package main

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
	return nil
}

func (v1) Receive() error {
	return nil
}

func main() {
	var rn ReceiveNotifier
	var r Receiver

	rn = v1{}

	// Можна робити приведення до інтерфейсу, так як rn реалізовує інтерфейс Receiver.
	r = rn

	// Хоч r містить тип v1, що реалізовує інтерфейс ReceiveNotifier.
	// Тип змінної Receiver, отже гарантувати, що вона містить реалізацію Notifier не можна.
	rn = r
}
