package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[2:3]

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("-------------------------------")

	// s2 посилається на той самий масив, отже при додаванні елементу, масив змінюється для обох зрізів
	fmt.Println("Add 10 and 20 to the s2, fits the capacity.")
	s2 = append(s2, 10)
	s2 = append(s2, 20)

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
	fmt.Println("-------------------------------")

	// Після того, як новий елемент не поміщається в існуючий масив, для s2 створюється новий масив
	fmt.Println("Add 30, overflows the capacity")
	s2 = append(s2, 30)

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}
