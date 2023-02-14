package main

import "fmt"

type S struct {
	Name string
}

func (s *S) Set(name string) {
	s.Name = name
}

func (s S) Get() string {
	return s.Name
}

func main() {
	t := S{}

	// Виклик методу як функції. ТАК РОБИТИ НЕ ТРЕБА, ЛИШЕ ДЛЯ ДЕМОНСТРАЦІЇ
	(*S).Set(&t, "test")

	fmt.Println(t)

	// Виклик методу як функції. ТАК РОБИТИ НЕ ТРЕБА, ЛИШЕ ДЛЯ ДЕМОНСТРАЦІЇ
	fmt.Println(S.Get(t))
}
