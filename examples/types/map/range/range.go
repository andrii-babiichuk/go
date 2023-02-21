package main

import "fmt"

func main() {
	// Мапа з визначеними простими числами
	prime := map[string]bool{
		"1": true,
		"2": true,
		"3": true,
		"4": false,
		"5": true,
		"6": false,
		"7": true,
		"8": false,
		"9": false,
	}

	// Прохід по елементах мапи не гарантує порядок.
	for k, v := range prime {
		fmt.Println(k, v)
	}
}
