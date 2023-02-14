// Приклад того, як порядок визначення полів в структурі впливає на розмір структури.

package main

import (
	"fmt"
	"unsafe"
)

type Type struct {
	float  float32
	int    int8
	double float64
	flag   bool
}

type Type2 struct {
	double float64
	float  float32
	int    int8
	flag   bool
}

func main() {
	var a Type
	// var a Type2

	fmt.Printf("float    address: %p, size: %d\n", &a.float, unsafe.Sizeof(a.float))
	fmt.Printf("int      address: %p, size: %d\n", &a.int, unsafe.Sizeof(a.int))
	fmt.Printf("double   address: %p, size: %d\n", &a.double, unsafe.Sizeof(a.double))
	fmt.Printf("flag     address: %p, size: %d\n", &a.flag, unsafe.Sizeof(a.flag))
	fmt.Println("-----------------------------------------")
	fmt.Printf("UserType address: %p, size: %d\n", &a, unsafe.Sizeof(a))
}
