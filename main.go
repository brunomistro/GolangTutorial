package main

import "fmt"

//PASS BY VALUE - PART 1

// func updateName(x string) {
// 	x = "yoshi"
// }

func updateName(x string) string {
	x = "wedge"
	return x
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func part1() {
	// group A types -> strings, ints, bools, floats, arrays, structs
	// non-pointer wrapper values
	name := "tifa"

	// updateName(name)
	name = updateName(name)

	fmt.Println(name)

	// group B types -> slices, maps, functions
	// pointer wrapper values
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu)
}

// POINTERS - PART 2

// func updateName(n string) {
// 	n = "wedge"
// }

func updateNameP(n *string) {
	*n = "wedge"
}

func part2() {
	name := "tifa"

	// updateNameP(name)
	// fmt.Println(name)

	// & gets the memory address of the value (pointer)
	fmt.Println("memory address of name is:", &name)

	// * gets the value at the specified memory address
	m := &name
	fmt.Println("memory address:", m)
	fmt.Println("value at memory address:", *m)

	updateNameP(m) // using pointer as arg, can pass &name directly instead of "m" as well
	fmt.Println(name)

	/*
	|--name---|----m----|
	|  0x001  |  0x002  |
	|---------|---------|
	| "tifa"  | p0x001  |
	|---------|---------|
	*/
}

func main() {
	part1()
	part2()
}
