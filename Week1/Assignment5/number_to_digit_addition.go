package main

import (
	"fmt" 
	"strconv"
)

func numberToDigitAddition(number int) {
	var numstr string
	numstr = strconv.Itoa(number)

	var arr []int

	for _,ch := range numstr {
		digit := int(ch - '0')
		arr = append(arr,digit)
	}
	result := 0
	for _,digit := range arr {

		result += digit
	}

	fmt.Println("This is the sum of the given number: ",result)
}

func main() {

	var number int
	fmt.Print("Enter Number: ")
	fmt.Scan(&number)

	numberToDigitAddition(number)
}