package main

import "fmt"

func foo() func() {
	return func(){
		fmt.Println("I am from the returns func")
	}
}

func main(){
	f := foo()
	f()
}
