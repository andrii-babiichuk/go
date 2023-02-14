package main

import (
	"errors"
	"fmt"
)

// ValidateString - функція може приймати 0 аб безліч аргументів та повертати 0 або безліч значень.
// При повернення 2 значень, як правло 2-ге це помилка.
// Більше 2-х значень або 2 значення без помилки не рекомендується використовувати, без вагомих причин.
func ValidateString(text string) (bool, error) {
	if len(text) > 0 {
		return true, nil
	}
	return false, errors.New("empty string")
}

// F1 - аргументи однакових типів можна об'єднувати, і писати тип лише раз.
func F1(arg1, arg2 int) {
	fmt.Printf("%T, %v\n", arg1, arg1)
	fmt.Printf("%T, %v\n", arg2, arg2)
}

// ValidateString2 - функція може мати іменовані змінні результату.
// При чому іменовані можуть бути лише всі значення.
func ValidateString2(text string) (valid bool, err error) {
	// Іменовані параметри можна розглядати як спрощений запис ініціалізації змінних на початку функції
	// var valid bool
	// var err error
	if len(text) > 0 {
		valid = true
	} else {
		err = errors.New("empty string")
	}
	// При іменованих аргументах можна повертати порожній return.
	// Такий варіант не рекомендується для складних функцій.
	return
	// return valid, err
}

// Sum - функція може мати змінну кількість аргументів одного типу.
// В такому випадку ці аргемути завжди мають іти в кінці і позначаються '...' перед типом.
func Sum(prefix string, numbers ...int) (a int) {
	fmt.Printf("%s: ", prefix)
	for i := 0; i < len(numbers); i++ {
		a += numbers[i]
	}
	return a
}

// Callback Функції це значення, їх можна отримувати як аргументи та повертати
func Callback(callback func() error) func() error {
	return func() error {
		fmt.Println("inner function")
		return callback()
	}
}

func main() {
	valid, err := ValidateString("e")
	fmt.Println(valid, err)

	// У випадку, якщо одне або кілька значень, що повертає функція не використовуються можна виконати присвоєння до "_".
	_, err = ValidateString2("")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("--------------")
	// Функції можна присвоювати
	a := func() error {
		fmt.Println("callback function")
		return nil
	}
	Callback(a)()

	fmt.Println("--------------")
	fmt.Println(Sum("sum", 1, 2, 3))
}
