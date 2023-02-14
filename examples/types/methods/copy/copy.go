package main

import "fmt"

type tt struct {
	Name string
}

func (t *tt) PrintPointer() {
	fmt.Println(t.Name)
}

func (t tt) PrintValue() {
	fmt.Println(t.Name)
}

func main() {

	a := tt{
		Name: "1",
	}

	// Копіюється посилання на функцію, та посилання на обʼєкт a
	pointer := a.PrintPointer

	// Копіюється посилання на функцію, та значення обʼєкту a
	value := a.PrintValue

	a.Name = "2"

	pointer()
	value()
}
