package main

import "fmt"

// User - новий тип визначається за допомогою ключового слова type.
// Як правило новий тип це структура, що може мати різні поля.
type User struct {
	Name string
	Age  int
}

func main() {
	// Нульове значення для типу буде структура з нульовими значеннями для кожного поля.
	var user User
	fmt.Printf("%+v\n", user)

	// Ключове слово new дозволяє створювати значення заданого типу, та повертає покажчик на значення.
	//user2 := new(User)
	//fmt.Printf("%+v\n", user2)

	//user3 := User{}
	//fmt.Printf("%+v\n", user3)

	// Визначення User з заданими полями, не обов'язково задавати всі поля.
	childUser := User{
		Age: 1,
	}
	fmt.Printf("%+v\n", childUser)

	// Визначення User опускаючи назву поля, в такому випадку порядок обовʼязковий, всі поля мають бути визначені.
	oldUser := User{"Old", 90}
	fmt.Printf("%+v\n", oldUser)
	fmt.Println("-----------")

	// Визначення не іменована структура
	b := struct {
		flag    bool
		counter int
	}{
		flag:    true,
		counter: 1,
	}
	fmt.Printf("%+v\n", b)
}
