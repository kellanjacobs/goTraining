package main

import "fmt"

func int_even(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func greatest(x ...float64) {
	var y float64
	for _, v := range x {
		if y == 0.0 {
			y = v
		}
		if v > y {
			y = v
		}
	}
	fmt.Println(y)
}

func main() {
	fmt.Println(int_even(3))
	fmt.Println(int_even(4))

	int_even_two := func(x int) (int, bool) {
		return x / 2, x%2 == 0
	}
	fmt.Println(int_even_two(5))
	fmt.Println(int_even_two(6))
	greatest(4, 3, 7, 10, 1)
}
