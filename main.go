package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getinput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	name, err := r.ReadString('\n')
	return strings.TrimSpace(name), err
}
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter Bill Name:")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	// b := newBill(name)

	name, _ := getinput("Enter Bill Name:", reader)

	b := newBill(name)
	fmt.Println("Created the bill named", b.name)
	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getinput("chose an option (a - add item | s - save bill | t - add tip)", reader)

	switch opt {
	case "a":
		name, _ := getinput("Item name", reader)
		price, _ := getinput("Item price", reader)
		fmt.Println(name, price)
	case "s":
		fmt.Println("you chose s")
	case "t":
		tip, _ := getinput("tip amount", reader)
		fmt.Println(tip)
	default:
		fmt.Println("that is not an option")
	}

	// fmt.Println(opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
