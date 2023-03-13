package main

import (
	"fmt"
)

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func sayGreeting(n string) {
	fmt.Printf("good morning %v \n", n)
}

func main() {

	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, ":", v)
	}
}
