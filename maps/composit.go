package main

import "fmt"

func main() {
	myNames := map[string]string{
		"Kellan": "Jacobs",
		"John":   "Doe",
	}
	myNames["Daniel"] = "Smith"
	fmt.Println(myNames)
	delete(myNames, "John")
	fmt.Println(myNames)
}
