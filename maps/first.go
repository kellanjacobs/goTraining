package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)

	myGreeting["Kellan"] = "Greetings"
	myGreeting["Jacobs"] = "Sawat Dii"
	fmt.Println(myGreeting)
}
