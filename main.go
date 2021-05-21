package main

import "fmt"

//Printing & Formatting Strings

func main() {
	age := 22
	name := "Bruno"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name)
	/*
		So, printing works pretty much as expected in print line and just print, \n works 2
		For getting variables and strings on a print, just separate and call the variables
	*/

	// Printf (formatted string), %_ = format specifier
	// always add \n !!!!
	fmt.Printf("my age is %v and my name is %v \n", age, name) // %v = value in default format
	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q = quotes
	fmt.Printf("age is of type %T \n", age)                    // %T is the type
	fmt.Printf("you scored %f points! \n", 225.55)             // %f = float format
	fmt.Printf("you scored %0.1f points! \n", 225.55)          // %0.2f = float with 2 decimal points

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
	//don't know why use this if can just print directly, but exists
}
