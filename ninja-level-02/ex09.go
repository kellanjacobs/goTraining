package main

import "fmt"

func main() {
	x := 2
	fmt.Printf("Dec: %d Binary: %b Hex: %#x\n", x, x, x)
	y := x << 1
	fmt.Printf("Dec: %d Binary: %b Hex: %#x\n", y, y, y)
}
