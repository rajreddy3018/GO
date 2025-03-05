package main

import "fmt"

func fibonacciNumber(number int) []int {

	a,b := 0,1

	var result []int
	for i:=0;i<number;i++ {
		result = append(result,a)
		a,b = b,a+b
	}

	return result
}

func main() {
	
	var number int

	fmt.Print("Enter n number to get fibonacci until that number:")
	fmt.Scan(&number)

	fmt.Printf("These are the fibonacci series of %d: %d\n", number, fibonacciNumber(number))
	
}