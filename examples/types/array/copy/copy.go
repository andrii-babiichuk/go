package main

import "fmt"

func main() {
	var a [3]string
	b := [3]string{"a", "b", "c"}

	// Виведення адрес елементів масиву b.
	for i := 0; i < len(b); i++ {
		fmt.Printf("%d: %p\n", i, &b[i])
	}

	fmt.Println("---------------")
	// При присвоєнні відбувається копіювання даних
	a = b

	// Виведення адрес елементів масиву a.
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d: %p\n", i, &a[i])
	}
}
