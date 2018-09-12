package main

import "fmt"

func main() {
	Greetings := map[int]string{
		0: "Good Morning",
		1: "Bonjour",
		2: "Buenos Dias",
		3: "Bongiorno",
	}
	fmt.Println(Greetings)

	delete(Greetings, 2)

	if val, exists := Greetings[2]; exists {
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	} else {
		fmt.Println("That Value doesnt exists")
		fmt.Println("Val: ", val)
		fmt.Println("Exists: ", exists)
	}
	fmt.Println(Greetings)
}
