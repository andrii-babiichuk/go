package main

import "fmt"

// Окрім структурних типів є можливість визначати нові типи на основі примітивів.
type ID string

func main() {
	var newId ID
	var id = "1"

	// Потребує явного приведення, хоч базуються на основі одного типу.
	newId = ID(id)

	fmt.Printf("%T\n", id)
	fmt.Printf("%T\n", newId)
}
