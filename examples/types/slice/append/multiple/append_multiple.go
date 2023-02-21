package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5}

	// Додавання до зрізу елементів з іншого зрізу
	a = append(a, b...)

	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))
	fmt.Println(a)
}
