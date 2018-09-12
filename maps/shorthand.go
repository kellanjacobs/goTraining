package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	myGreeting["Kellan"] = "Howdy"
	myGreeting["Jacobs"] = "Hello"
	fmt.Println(myGreeting)
}
