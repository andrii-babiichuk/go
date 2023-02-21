package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Визначення зрізу з нульовим значенням
	var zero []int
	// Визначення порожнього зрізу
	empty := []int{}
	empty2 := make([]int, 0)

	fmt.Println("Name   | Value | Size | nil   | len | cap | Pointer to backing array")
	fmt.Printf("Zero   | %v    | %d   | %t  | %d   | %d   | %p \n", zero, unsafe.Sizeof(zero), zero == nil, len(zero), cap(zero), zero)
	fmt.Printf("Empty  | %v    | %d   | %t | %d   | %d   | %p \n", empty, unsafe.Sizeof(empty), empty == nil, len(empty), cap(empty), empty)
	fmt.Printf("Empty2 | %v    | %d   | %t | %d   | %d   | %p \n", empty2, unsafe.Sizeof(empty2), empty2 == nil, len(empty2), cap(empty2), empty2)

	// Всі методи роботи зі зрізом з нульовим значенням валідні, і не викликають помилки.
	for i, v := range zero {
		fmt.Println(i, v)
	}
	d := zero[:]
	fmt.Println(d)

	// Важлива різниця лише при роботі з JSON, зріз nil конвертується в "null", тоді як порожній зріз в "[]".
}
