package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{20, 40, 30}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 1)
	fmt.Print(index)
}
