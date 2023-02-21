package main

import "fmt"

// Копіюється весь масив
func last(arr [1e6]int) {
	fmt.Printf("%p\n", &arr)
}

// Копіюється посилання на масив
func last2(arr *[1e6]int) {
	fmt.Printf("%p\n", arr)
}

func main() {
	var a [1e6]int

	fmt.Printf("%p\n", &a)
	last(a)
	last2(&a)
}
