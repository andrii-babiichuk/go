package main

import "fmt"

type Task struct {
	Name     string
	IsDone   bool
	Assigner string
}

// Визначення методу до покажчика (pointer receiver)
func (t *Task) Do() {
	t.IsDone = true
}

// Визначення методу до значення (value receiver)
func (t Task) Assign(assigner string) {
	t.Assigner = assigner
}

func main() {
	t := Task{
		Name: "First",
	}
	fmt.Printf("%+v\n", t)

	// (&t).Do() // Go виконує автоматичне приведення до покажчика і навпаки
	t.Do()
	fmt.Printf("%+v\n", t)

	p := &t
	p.Assign("me")
	// (*p).Assign() // Go виконує автоматичне приведення до покажчика і навпаки
	fmt.Printf("%+v\n", t)
}
