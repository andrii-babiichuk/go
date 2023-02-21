package main

import "fmt"

func main() {
	a := [3]int{10, 20, 30}

	fmt.Println(len(a)) // довжина масиву

	// Прохід масиву через for
	// При такому проходженні відбувається робота з оригінальними значеннями масиву
	for i := 0; i < len(a); i++ {
		a[i] = a[i] * 2
		fmt.Printf("%d: value: %d, address a[i]: %p\n", i, a[i], &a[i])
	}
	fmt.Println(a)
	fmt.Println("----------")

	b := [3]int{10, 20, 30}
	// Range повертає 2 значення: індекс та копію значення.
	// В такому випадку зміна значення v не змінить сам масив.
	// Якщо потрібна мутація можна використовувати b[i] так само як і в класичному for
	for i, v := range b {
		v = v * 2
		fmt.Printf("%d: value: %d, address v: %p\n", i, v, &v)
	}
	fmt.Println(b)

	//fmt.Println("----------")
	//// range дозволяє використовувати лише індекс, тоді значення v не буде копіюватись.
	//for i := range b {
	//	fmt.Println(b[i])
	//}

	// У випадку, якщо індекс не потрібен, його можна опускати.
	//for _, v := range b {
	//	fmt.Println(v)
	//}
}
