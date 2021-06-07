package main

import "fmt"

//Receiver Functions with Pointers

func main() {
	mybill := newBill("Bruno")
	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 8.95)
	mybill.addItem("toffee pudding", 4.95)
	mybill.addItem("coffee", 3.25)
	mybill.updateTip(10)
	fmt.Println(mybill.format())
}
