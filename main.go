package main

import "fmt"

func main() {
	// var nameOne = "one"
	// var nameTwo = "Two"
	// nameOne = "three"
	// nameThree := "four"
	var name string = "ahmed"
	var age int = 24
	var forrmated = fmt.Sprintf("My name is %v and my age is %v", name, age)
	fmt.Println(forrmated)
	// fmt.Println(nameOne, nameTwo, nameThree)

	var ages = [3]int{1, 2, 3}

	fmt.Println(ages, len(ages))
}
