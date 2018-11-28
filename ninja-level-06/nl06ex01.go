package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 451
}

func bar() (string, int){
	return "Matt Smith", 75
}
