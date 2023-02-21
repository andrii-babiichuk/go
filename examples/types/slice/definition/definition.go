package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a []int            // Ініціалізація порожнього зрізу
	b := []int{1, 2, 3}    // Спрощена форма
	c := make([]int, 3)    // Створення зрізу довжиною і місткістю 3, за допомогою вбудованої функції make
	d := make([]int, 3, 5) // Створення зрізу довжиною 3 і місткістю 5, за допомогою вбудованої функції make

	fmt.Println(a, b, c, d)
	fmt.Println(unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c), unsafe.Sizeof(d))
}
