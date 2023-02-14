// Приклад приведення типів між іменованими структурами та анонімною структурою, що мають однаковий формат.

package main

import "fmt"

type UserType struct {
	Name    string
	Counter int
}

type UserType2 struct {
	Name    string
	Counter int
}

func main() {
	var a UserType
	var b UserType2

	// Анонімний тип, що має такий же формат, як і UserType/UserType2.
	c := struct {
		Name    string
		Counter int
	}{
		Name:    "name",
		Counter: 1,
	}

	// Неявне приведення типів між двома різними структурами не дозволяється.
	// a = b // cannot use b (variable of type UserType2) as UserType value in assignment
	a = UserType(b)

	// Неявне приведення типів між іменованою структурою та анонімною дозволяється.
	b = c
	c = a

	fmt.Printf("%T: %+v\n", a, a)
	fmt.Printf("%T: %+v\n", b, b)
	fmt.Printf("%T: %+v\n", c, c)
}

// Анонімні структури можна використовувати як аргументи
func fullInfo(s struct {
	Name    string
	Counter int
}) string {
	return fmt.Sprintf("Name: %s, Age: %d", s.Name, s.Counter)
}
