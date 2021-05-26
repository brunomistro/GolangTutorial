package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//The Standard Library - Part 1
	part1()
	//Loops - Part 2
	part2()
	//Booleans & Conditionals - Part 3
	part3()
}

func part1() {
	fmt.Println("Begining of The Standard Library - Part 1")

	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.Contains(greeting, "bye"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "bowser"))

	fmt.Println("## End of The Standard Library - Part 1 ##")
	fmt.Println("")
}

func part2() {
	fmt.Println("Begining of Loops - Part 2")

	x := 0
	for x < 5 {
		fmt.Println("value of x is:", x)
		x++ // infinite loop without this
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, val := range names {
		fmt.Printf("the value at pos %v is %v \n", index, val)
		val = "new string"
	}

	for _, val := range names {
		fmt.Print(val, ",")
		val = "new string"
	}

	// changing val in a loop does not mutate the original slice
	fmt.Println(names)

	fmt.Println("## End of Loops - Part 2 ##")
	fmt.Println("")
}

func part3() {
	fmt.Println("Begining of Booleans & Conditionals - Part 3")

	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 40")
	}

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, val := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v \n", index, val)
	}

	fmt.Println("## End of Booleans & Conditionals - Part 3 ##")
	fmt.Println("")
}
