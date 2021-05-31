package main

import (
	"fmt"
	"strings"
)

//Multiple Return Values - Part 1

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
		/*
			to get the last element of a slice is array[len(array)-1]
		*/
	}

	return initials[0], "_"
}

func multReturnValues() {
	fn1, sn1 := getInitials("tifa lockhart")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInitials("cloud strife")
	fmt.Println(fn2, sn2)

	fn3, sn3 := getInitials("barret")
	fmt.Println(fn3, sn3)
}

//Package Scope - Part 2

func packageScope() {
	sayHello("mario")
	showScore()

	for _, v := range points {
		fmt.Println(v)
	}
}

var score float64 = 99.5

func main() {
	//score = 99.5
	multReturnValues()
	packageScope()
}
