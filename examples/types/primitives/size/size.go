package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Числові типи можуть бути фіксованого розміру, задається як кількість біт біля типу (int64, uint8, float64)
	var fixedSize int32

	// Розмір типів int та uint залежить від архітектури. В 32-х бітній архітектурі вони займають 4 байти, в 64-х бітній - 8.
	// “Розмір слова” (word size) - обсяг пам'яті, необхідний для зберігання цілого числа та покажчика для даної архітектури.
	var variableSize int

	// Строки будуються на основі масиву байт.
	// Розмір строкового типу складає 2 слова, перше містить посилання на масив, друге - його розмір.
	var empty string

	// unsafe.Sizeof - повертає розмір типу, проте не включає будь-які посилання на додаткові структури.
	// Тому розмір будь-якої строки повертатиметься, як 2 слова.
	var defined = "some text"

	// Булеан - 1 байт.
	var c bool

	fmt.Printf("%T: %d\n", fixedSize, unsafe.Sizeof(fixedSize))
	fmt.Printf("%T: %d\n", variableSize, unsafe.Sizeof(variableSize))
	fmt.Printf("%T: %d\n", empty, unsafe.Sizeof(empty))
	fmt.Printf("%T: %d\n", defined, unsafe.Sizeof(defined))
	fmt.Printf("%T: %d\n", c, unsafe.Sizeof(c))
}
