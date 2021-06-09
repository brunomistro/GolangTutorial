package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Parsing Floats

// func getInput(prompt string, r *bufio.Reader) (string, error) {
// 	fmt.Print(prompt)
// 	input, err := r.ReadString('\n')
// 	return strings.TrimSpace(input), err
// }

func getInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	//reader := bufio.NewReader(os.Stdin)
	/*
		Trying to do the reader on func, works properly,
		I'll let you know in future branches,
		if that has any problems
		(The video func above)
	*/
	name, _ := getInput("Create a new bill name: ")
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOptions(b bill) {
	opt, _ := getInput("Choose option (a -add item, s - save bill, t - add tip): ")
	switch opt {
	case "a":
		name, _ := getInput("Item name: ")
		price, _ := getInput("Item price: ")

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number!")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("Item added - ", name, price)
		promptOptions(b)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ")

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number!")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("Tip added - ", tip)
		promptOptions(b)

	case "s":
		fmt.Println("you chose to save the bill", b)

	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
