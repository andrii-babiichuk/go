package main

import "fmt"

func main() {
	var a []int
	// a := make([]int, 0, 100000)

	capacity := cap(a)
	for i := 0; i < 100000; i++ {
		a = append(a, i*i)
		if cap(a) != capacity {
			if capacity != 0 {
				fmt.Printf("New array address: %p, cap: %d, inc: +%d%%\n", a, cap(a), (cap(a)*100)/capacity-100)
			}
			capacity = cap(a)
		}
	}

	fmt.Printf("len: %d, cap: %d\n", len(a), cap(a))
}
