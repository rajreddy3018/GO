package main

import "fmt"

func SimpleInterest(principal int, rate float64, time int) {

	var result float64

	result = (float64(principal) * rate * float64(time)) / 100

	fmt.Println("This is Simple Interst Value: ", result)

}

func main() {
	var principal int
	var rate float64
	var time int

	fmt.Print("Enter Principal Amount: ")
	fmt.Scan(&principal)
	fmt.Print("Enter Rate in percentage: ")
	fmt.Scan(&rate)
	fmt.Print("Enter time in years: ")
	fmt.Scan(&time)

	SimpleInterest(principal, rate, time)

}