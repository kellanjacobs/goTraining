package main

import "fmt"

type person struct {
	first string
}
func changeMe(p *person) {
	p.first = "Kellan"
}

func main() {
	p1 := person{"Daniel"}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}