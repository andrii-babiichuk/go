package pack

import "fmt"

// Тип unexported не експортується за межі пакету.
type unexported struct {
	// Поле доступне в інших пакетах.
	Name string
	// Поле не доступне в інших пакетах.
	secret int
}

// Непублічний метод
func (u unexported) print() {
	fmt.Print(u)
}

// Публічний метод
func (u unexported) Println() {
	u.print()
}

// NewUnexported доступний в інших пакетах, проте тип unexported ні.
// Такої ситуації варто уникати. Всі методи, що доступні публічно мають повертати публічні типи.
func NewUnexported() unexported {
	return unexported{"Test", 1}
}
