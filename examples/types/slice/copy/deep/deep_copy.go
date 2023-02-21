package main

import "fmt"

func main() {
	destination := make([]string, 3)
	source := []string{"a", "b", "c"}

	// Функція copy копіює значення зрізу в інший. Але при цьому не додає і не видаляє елементи.
	fmt.Printf("Copied %d elements from source to the destination.\n", copy(destination, source))

	source[0] = "A"

	fmt.Println("-----------------")
	// Виведення адрес елементів масиву a
	fmt.Println("Source slice:")
	for i := range source {
		fmt.Printf("%d: %p, %s\n", i, &source[i], source[i])
	}
	fmt.Println("-----------------")
	// Виведення адрес елементів масиву b
	fmt.Println("Destination slice:")
	for i := range destination {
		fmt.Printf("%d: %p, %s\n", i, &destination[i], destination[i])
	}
}
