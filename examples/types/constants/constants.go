package main

import (
	"fmt"
)

// Визначення константи
const koef = 2

// Константи можуть бути лише примітивами
const (
	cBool   = true
	cString = "Types"
	cNumber = 2
)

// Константи можна визначати за допомогою арифметичних чи бітових операцій зі значеннями та іншими константами.
const sq = koef * 2

// Але не можна використовувати функції чи змінні
// const sq2 = math.Sin(30)

func n(a int, b float32) {
	fmt.Println(a, b)
}

func main() {
	i := 1   // int
	f := 1.1 // float64

	// Константи це значення, а не змінні.
	// Числові константи в залежності від використання можуть "змінювати тип"
	fmt.Println(i * koef)
	fmt.Println(f * koef)

	n(koef, koef)
}
