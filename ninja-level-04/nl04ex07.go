package main

import "fmt"

func main() {
	jb := []string{"james", "bond", "Shaken, not stirred",}
	mp := []string{"Miss", "Moneypenny", "Helloooooo, James.",}
	people := [][]string{jb,mp}
	fmt.Println(people)
	for i, person := range people {
		fmt.Println("Record #", i)
		for j, val := range person{
			fmt.Printf("\tIndex: %v\t Value: %s\n", j,val)
		}
	}
}
