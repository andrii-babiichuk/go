package main

import "fmt"

type N struct {
	Name string
}

func (t *N) PrintlnPointer() {
	fmt.Println("pointer")
}

func (t N) PrintValue() {
	fmt.Println("value")
}

func main() {
	// Визначення змінної зі значенням nil
	var t *N
	fmt.Println(t)

	// Виклик методу, що працює з покажчиком
	t.PrintlnPointer()

	// Виклик методу, що працює зі значенням
	// t.PrintValue()
}
