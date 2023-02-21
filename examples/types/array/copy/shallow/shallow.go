package main

import "fmt"

func main() {
	// Визначення масиву з покажчиками на строки
	var c [3]*string
	d := [3]*string{new(string), new(string), new(string)}

	// Присвоєння значень
	*d[0] = "a"
	*d[1] = "b"
	*d[2] = "c"

	// Копіювання покажчиків на значення
	c = d

	*c[0] = "d"

	// Виведення значень елементів масиву a
	for i := 0; i < len(c); i++ {
		fmt.Printf("[%d] address: %p, value: %s \n", i, c[i], *c[i])
	}

	fmt.Println("--------------")
	// Виведення значень елементів масиву b
	for i := 0; i < len(d); i++ {
		fmt.Printf("[%d] address: %p, value: %s \n", i, d[i], *d[i])
	}
}
