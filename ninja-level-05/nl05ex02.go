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
		last: "Somvong",
		fav: []string{"Chocolate",
			"Coffee",},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v.first, v.last, v.fav)
	}
}
