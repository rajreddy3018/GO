package main

import "fmt"


func addingNumber(number1 float64, number2 float64) {
	result:= number1+number2
	fmt.Println("This is the addition of ",number1, "and ", number2, " : ", result)
}

func subtractingNumber(number1 float64, number2 float64) {
	result:= number1-number2
	fmt.Println("This is the subtraction of ",number1, "and ", number2, " : ", result)
}

func multiplyingNumber(number1 float64, number2 float64) {
	result:= number1*number2
	fmt.Println("This is the multiplication of ",number1, "and ", number2, " : ", result)
}

func dividingNumber(number1 float64, number2 float64) {
	result:= number1/number2
	fmt.Println("This is the division of ",number1, "and ", number2, " : ", result)
}



func main() {
	var number1 float64
	var  number2 float64
	var operation string

	fmt.Print("Enter First Number: ")
	fmt.Scan(&number1)
	fmt.Print("Enter Second Number: ")
	fmt.Scan(&number2)
	fmt.Print("Enter Symbol to Perform Operations: ")
	fmt.Scan(&operation)
	switch operation {
		case "+": addingNumber(number1, number2)
		case "-": subtractingNumber(number1, number2)
		case "*": multiplyingNumber(number1, number2)
		case "/": dividingNumber(number1, number2)
		default: fmt.Println("Select the Operators only from these: +, - , *, / ")
	}
}