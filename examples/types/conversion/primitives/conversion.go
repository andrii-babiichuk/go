package main

import (
	"fmt"
	"strconv"
)

func main() {
	number := 3
	text := "2"
	bolean := true

	// number = int(bolean) // несумісні типи не можливо перетворити
	// коректний запис буде
	if bolean {
		number = 1
	}
	// і навпаки
	bolean = number != 0

	// конвертацію строк можна виконувати за допомогою пакету strconv
	number, _ = strconv.Atoi(text)

	text = strconv.Itoa(number)

	// До строки будь-який тип можна привести за допомогою функції Sprint/Sprintf
	text = fmt.Sprint(bolean)

	// Конвертація до float. Зверніть увагу, на роботу ParseFloat.
	// Якщо використовується bitSize: 32, значення потрібно явно приводити до float32. Інакше точність втратиться.
	floatAsString := "2.4"
	float1, _ := strconv.ParseFloat(floatAsString, 32)
	float2, _ := strconv.ParseFloat(floatAsString, 64)
	fmt.Println(float1)
	fmt.Println(float2)
	fmt.Println(float32(float1))
	fmt.Println(float32(float2))
}
