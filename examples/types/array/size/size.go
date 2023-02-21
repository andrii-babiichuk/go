package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Порожній масив з елементами типу int
	var a [4]int
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println("------------------")

	//// Розмір масиву не можна визначити на основі змінної, можна на основі константи
	//var size = 10
	//var variableSize [size]int
	//fmt.Println(variableSize)

	var c = [4]string{"A", "B", "C", "DEF"}
	fmt.Println(unsafe.Sizeof(c))
}
