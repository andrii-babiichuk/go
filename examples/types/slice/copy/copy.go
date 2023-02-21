package main

import "fmt"

func main() {
	destination := make([]string, 0)
	source := []string{"a", "b", "c"}

	// При присвоєнні відбувається копіювання даних.
	destination = source
	fmt.Println("Pointer to the backing array:")
	fmt.Printf("source:       %p\n", source)
	fmt.Printf("desctination: %p\n", destination)
	fmt.Println("-----------------------------------------------")

	// Відповідно зміна значення в одному приводить до зміни значення в іншому.
	source[0] = "A"
	fmt.Println("Zero element    | Pointer to the zero element")
	fmt.Printf("Source:       %s | %p\n", source[0], &source[0])
	fmt.Printf("Desctination: %s | %p\n", destination[0], &destination[0])
}
