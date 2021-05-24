package main

import "fmt"

//Arrays and Slices

func main() {

	// var ages [3]int = [3]int{20, 25, 30}
	/*
		Both declaring are the same, but it's easier to do 2nd version (below),
		'cause de 'int' after the '=' is mandatory,
		so no need to give it a type b4.
		In Go array are fixed length.
		START WITH 0!
	*/
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"} //short declare
	names[1] = "luigi"                                      //updating

	fmt.Println(ages, len(ages)) //'len()' length function
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	// and can be added values to that slice
	// NO NEEDED AMOUNT ON THE BRACKETS
	var scores = []int{100, 50, 60}
	scores[2] = 25 // updating
	scores = append(scores, 85)
	/* So, just using 'append()', it's going to return the scores slice whit a value, 85 in this case,
	then for adding a value to the slice, u need to equals to the slice, as above.
	*/

	fmt.Println(scores, len(scores))

	// slice ranges , kinda substrings ?? dk
	// wierd AF
	rangeOne := names[1:4]  // doesn't include pos 4 element
	rangeTwo := names[2:]   //includes the last element
	rangeThree := names[:3] //does not includes the pos 3 element
	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}
