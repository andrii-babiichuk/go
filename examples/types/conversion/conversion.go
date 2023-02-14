package main

import "fmt"

func main() {
	var a int
	var b int32
	var c int64

	// При роботі з подібними типами всі перетворення відбуваються явно.
	// В даному випадку дані копіюються з новим типом.
	a = int(b)
	c = int64(a)
	b = int32(c)

	// Всі наступні приведення будуть викликати помилку, хоча в залежності від архітектури
	// int це те саме, що int64 або int32
	//a = b
	//a = c
	//b = a
	//c = a

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	// Строку можна привести до зрізу байт і навпаки.
	fmt.Println("--------------")
	str := "Abc"
	sInBytes := []byte(str)

	fmt.Println(sInBytes)
}
