package main

import (
	"fmt"
)

func main() {
	// x := 0
	// for x < 5 {
	// 	fmt.Println(x)
	// 	x++
	// }
	names := []string{"ahmed", "ali", "khalid"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println("the value of i:", names[i])
	// }

	for index, value := range names {
		fmt.Printf("the postion index %v and the value is %v \n", index, value)
	}
}
