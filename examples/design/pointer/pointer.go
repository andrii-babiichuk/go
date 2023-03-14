package main

import "fmt"

func main() {
	// Визначення покажчика змінної типу number, що використовує семантику покажчика
	var n *number
	// Визначення покажчика змінної типу number, що використовує семантику значення
	var n2 *number2

	fmt.Println(n.toString())
	fmt.Println(n2.toString())

	// Створення значень змінних
	// fmt.Println(number(1).toString())
	// fmt.Println(number2(1).toString())

}

type number int

func (n *number) toString() string {
	return fmt.Sprintf("%d", n)
}

type number2 int

func (n number2) toString() string {
	if &n == nil {
		return fmt.Sprintf("invalid")
	}
	return fmt.Sprintf("%d", n)
}
