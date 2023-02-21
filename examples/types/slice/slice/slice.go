package main

import "fmt"

func main() {
	// Початковий зріз
	a := []int{1, 2, 3, 4, 5}

	// Від елементу з індексом 2 [включно] до кінця
	b := a[2:]

	// Від початку до елементу з індексом 3 [виключно]
	c := a[:3]

	// Від елементу з індексом 1 [включно] до елементу з індексом 3 [виключно]
	d := a[1:3]

	// Від елементу з індексом 1 [включно] до елементу з індексом 3 [виключно].
	// Та з Обмеження місткості до елементу з індексом 4 [виключно]
	e := a[1:3:4]

	fmt.Println("Slice   | len | cap | values")
	fmt.Printf("init    | %d   | %d   | %v\n", len(a), cap(a), a)
	fmt.Printf("[2:]    | %d   | %d   | %v\n", len(b), cap(b), b)
	fmt.Printf("[:3]    | %d   | %d   | %v\n", len(c), cap(c), c)
	fmt.Printf("[1:3]   | %d   | %d   | %v\n", len(d), cap(d), d)
	fmt.Printf("[1:3:3] | %d   | %d   | %v\n", len(e), cap(e), e)
	fmt.Println("-------------------")
	// Зрізи посилаються на один масив, зміна значення змінить і інші зрізи
	d[1] = 10

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println("-------------------")
	arr := [3]int{1, 2, 3}
	slice := arr[1:2] // створення зрізу на основі існуючого масиву

	fmt.Println(slice)
}
