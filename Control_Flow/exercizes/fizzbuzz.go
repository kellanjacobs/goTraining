package main

import "fmt"
import "github.com/fatih/color"

func main() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			color.Red("FizzBuzz")
		} else if i%3 == 0 {
			color.Blue("Fizz")
		} else if i%5 == 0 {
			color.Green("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
