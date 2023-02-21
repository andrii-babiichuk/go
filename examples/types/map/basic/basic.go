package main

import "fmt"

func main() {
	m3 := map[string]byte{"1": 1, "2": 4, "3": 9}
	m4 := make(map[string]string, 1024)

	// Характеристики мапи
	fmt.Printf("map length: %d\n", len(m4))

	// Основні операції з мапами.
	m3["0"] = 0
	fmt.Printf("m3[\"4\"]: %d\n", m3["4"])
	delete(m3, "1")

	fmt.Println(m3)

	// При отриманні значення по ключу, якщо ключа не має - повертається нульове значення типу.
	zero := m3["0"]
	one := m3["1"]

	fmt.Printf("zero value: %d, one value: %d\n", zero, one)

	// Для визначення, що значення не задано є другий параметр
	_, zeroExists := m3["0"]
	_, oneExists := m3["1"]

	fmt.Printf("zero value exists: %t, one value exists: %t", zeroExists, oneExists)
}
