package main

import "fmt"

type person struct{
	first string
	last string
	fav []string
}

func main() {
	p1 := person{
		first: "Daniel",
		last: "Nowak",
		fav: []string{"strawberry",
			          "Javachip",
			          "Earl Grey",},
	}
	p2 := person{
		first: "Tawan",
		last: "Nowak",
		fav: []string{"Chocolate",
					  "Coffee",},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.fav {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.fav {
		fmt.Println(i, v)
	}
}
