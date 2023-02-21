package main

import (
	"fmt"
)

func main() {
	a1 := [3]string{"12", "2", "3"}
	a2 := [3]string{"12", "2", "3"}

	var c1 [3]int
	c2 := [3]int{0, 0, 0}

	// Оператор порівняння порівнює значення масивів, а не посилання
	fmt.Println(a1 == a2)
	fmt.Println(c1 == c2)
}
