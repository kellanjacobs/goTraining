package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func (p person) speak(){
	fmt.Printf("My Name is %s %s\n", p.first, p.last)
	fmt.Printf("I am %v years old\n", p.age)
}
func main(){
	p1 := person{
		first: "Daniel",
		last: "Nowak",
		age: 43,
	}
	p1.speak()
}