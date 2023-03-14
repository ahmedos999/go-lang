package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")
	fmt.Println(mybill)

	mybill.updateTip(10.0)
	mybill.addItem("mojito", 20.99)

	fmt.Println(mybill.format())
}
