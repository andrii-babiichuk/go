package main

import (
	"fmt"
)

func main() {
	a := []string{"12", "2", "3"}
	b := []string{"12", "2", "3"}

	fmt.Println(a)
	fmt.Println(b)

	// Тип зріз не підтримує оператор порівняння
	// fmt.Println(a == b) // помилка
}
