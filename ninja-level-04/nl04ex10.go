package main

import "fmt"

func main(){
	m := map[string][]string {
		"bond_james": []string{`Shaken, not stirred`,
			`Martinis`,
			`Women`},
		"moneypenny_miss" : []string{`James Bond`,
			`Literature`,
			`Computer Science`},
		`no_dr`: []string{`Being evil`,
			`Ice cream`,
			`Sunsets`},

	}

	m["Daniel"] = []string{"Asian Twinks", "White Twinks", "All Twinks"}

	delete(m, "bond_james")
	delete(m, "moneypenny_miss")

	for k,v := range m {
		fmt.Println("Record #", k)
		for i, val := range v{
			fmt.Println("\t", i, val)
		}
	}
}
