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

	names := []string{"ahmed", "ali", "khalid"}
	cycleNames(names, sayGreeting)
}
