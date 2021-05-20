package main

import "fmt"

//Variables, Strings & Numbers

func main() {
	//Strings
	var nameOne string = "mario" //default
	var nameTwo = "luigui"       //type can be hidden, but still a string
	var nameThree string         //initialize as empty

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "browser"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "yoshi" //most used ':=' replace 'var' and type, but can only be used inside of a function
	fmt.Println(nameFour)

	// int variables
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Print(ageOne, ageTwo, ageThree)

	// bits & memory
	// var numOne int8 = 25
	// var numTwo int8 = 128 // too large a number for 8-bit
	// var numTwo uint = -25 unsigned ints cannot be negative

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 1965385877.5
	var scoreThree = 1.5 // inferred as float64

	fmt.Print(scoreOne, scoreTwo, scoreThree)

	//https://golang.org/ref/spec#Numeric_types
}
