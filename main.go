package main

import "fmt"

func main() {
	mybill := newBill("Bruno")
	fmt.Println(mybill.format())
}
