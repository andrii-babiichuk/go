package main

import "fmt"

// Коли тип з яки буде іти роботи може бути різний, використовується спеціальний тип interface{}.
// Тоді через конструкцію `switch v.(type) {` можна визначити який тип має значення.
func isZero(v interface{}) bool {
	switch v.(type) {
	case string:
		return v == ""
	case bool:
		// Зверніть увагу, що тип v, в даному випадку `interface{}`, а не `bool`.
		// Відповідно для того, щоб виконувати операції властиві типу `bool`, потрібно виконати приведення
		// return !v // invalid operation: operator ! not defined on v (variable of type interface{})
		// В такому випадку варто перевіряти чи тип коректний.
		// b, ok := !v.(bool)
		return !v.(bool)
	case int:
		return v == 0
	default:
		return v == nil
	}
}

func main() {
	var a = 1
	var b = ""
	var c = true

	fmt.Println(isZero(a))
	fmt.Println(isZero(b))
	fmt.Println(isZero(c))
}
