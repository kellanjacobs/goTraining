package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println("Called First")
}
func bar() {
	fmt.Println("Called Last")
}