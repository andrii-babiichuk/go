package main

import (
	"fmt"
)

// При ініціалізації змінних, якщо значення не задано явно, їм буде присвоєно нульове значення.
// Це означає, що компілятор зарезервує необхідний об`єм памʼяті, та визначить кожен байт, як нуль.
// Для різних типів нульове значення буде різне.

func main() {
	var v1 int    // 0
	var v2 bool   // false
	var v3 string // ""

	fmt.Println("zero values: ")
	fmt.Printf("v1 type: %T, value: %v\n", v1, v1)
	fmt.Printf("v2 type: %T, value: %v\n", v2, v2)
	fmt.Printf("v3 type: %T, value: %v\n", v3, v3)
}
