package main

import "demo/pack"

func main() {
	// Можливо створити змінну з неекспортованим типом.
	u := pack.NewUnexported()

	// Неможливо викликати неекспортований метод.
	// u.print()
	u.Println()
}
