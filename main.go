package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := "hello everybody"
	test := "test string"
	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "hey"))
	fmt.Println("this is to check", test, "how well it did")
}
