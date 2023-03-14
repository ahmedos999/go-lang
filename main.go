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

	fmt.Println(opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
