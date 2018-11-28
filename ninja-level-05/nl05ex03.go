package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	mytruck := truck{
		vehicle: vehicle{
			doors: 4,
			color: "Green",
		},
		fourWheel: false,
	}

	mycar := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Purple",
		},
		luxury: true,
	}

	fmt.Println(mytruck)
	fmt.Println(mycar)
	fmt.Println(mytruck.color)
	fmt.Println(mycar.color)
}
