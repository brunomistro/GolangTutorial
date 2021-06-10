package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Saving Files

func getInput(prompt string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
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
		b.save()
		fmt.Println("Bill has been saved as: ", b.name)

	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
