package main

import "fmt"

func EvenOrOdd(number int) {
	var result int
	result = number%2
	if (result == 0) {
		fmt.Println("This is a Even Number")
	} else {
		fmt.Println("This is a Odd Number")
	}
}

func main() {

	var number int

	fmt.Print("Enter A Number: ")
	fmt.Scan(&number)

	EvenOrOdd(number)

}