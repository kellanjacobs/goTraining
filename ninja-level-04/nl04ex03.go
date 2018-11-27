package main

import "fmt"


func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := x[:5]
	z := x[5:]
	t := x[2:7]
	s := x[1:6]
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(t)
	fmt.Println(s)
}

