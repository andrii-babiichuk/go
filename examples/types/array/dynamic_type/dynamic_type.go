package main

import "fmt"

func main() {
	// Використовуючи interface{} можна створити масив з елементами різних типів.
	var b [4]interface{}
	b[0] = "test"
	b[1] = int8(1)
	b[2] = true
	b[3] = [2]int{}

	fmt.Println(b)
}
