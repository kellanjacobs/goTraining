package main

import "fmt"

func main() {
	i := 0
	for i < 10 {
		i++
		if i%2 == 1 {
			continue
		}
		fmt.Println(i)

	}
}
