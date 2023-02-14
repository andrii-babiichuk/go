package main

import "fmt"

// Ідентифікатор iota представляє собою інкрементальне значення int.
// iota збільшується на 1 в межах блоку const.
// При визначенні нового блоку значення обнуляється.
// Як правило застосовується для як заміна enum.

// Приклад застосування як альтернатива enum для статусу.
const (
	Init       = iota
	InProgress // = iota
	Completed  // = iota
	Failed     // = iota
)

// Приклад застосування як альтернатива enum для місткості пам'яті.
const (
	B = 1 << (10 * iota)
	KB
	MB
	GB
)

func main() {
	size := 124243215345

	gb := size / GB
	mb := (size % GB) / MB
	fmt.Printf("Capacity is %d GB and %d MB", gb, mb)
}
