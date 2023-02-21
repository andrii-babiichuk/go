package main

import (
	"fmt"
)

func main() {
	a := "ІI"

	fmt.Println(len(a))
	fmt.Println("--------------")
	for i := range a {
		fmt.Printf("%d: %c %b\n", i, a[i], a[i])
	}
	//fmt.Println("--------------")
	//for i := 0; i < len(a); i++ {
	//	fmt.Printf("%d: %c %b\n", i, a[i], a[i])
	//}
	//fmt.Println("--------------")
	for i, v := range a {
		// fmt.Println(utf8.RuneLen(a[i]))
		// fmt.Println(utf8.RuneLen(v))
		fmt.Printf("%d: %c %b\n", i, v, v)
	}
}
