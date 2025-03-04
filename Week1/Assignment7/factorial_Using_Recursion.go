package main

import "fmt"

func factorialValue(factorialNumber int) int {

	if (factorialNumber == 0 || factorialNumber == 1) {
		return 1
	}
	return factorialNumber * factorialValue(factorialNumber - 1)

}

func main() {

	var factorialNumber int

	fmt.Print("Enter the factorial number value: ")
	fmt.Scan(&factorialNumber)

	fmt.Printf("This is the value of factorial %d is %d\n",factorialNumber, factorialValue(factorialNumber))
}