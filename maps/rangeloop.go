package main

import "fmt"

func main() {
	Greetings := map[int]string{
		0: "Good Morning",
		1: "Bonjour",
		2: "Buenos Dias",
		3: "Bongiorno",
	}

	for key, val := range Greetings {
		fmt.Println(key, " - ", val)
	}
}
