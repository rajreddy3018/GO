package main

import "fmt"

func divisionOperation(dv1 int, dv2 int) {

	if (dv2 == 0) {
		fmt.Println("Denominator cannot be zero")
	} else {
		result:=dv1/dv2
		fmt.Printf("This is division value of %d and %d: %d\n", dv1, dv2, result)
	}
	

}

func main() {
	var divisionNumber1 int
	var divisionNumber2 int
	fmt.Print("Enter First Division Number: ")
	fmt.Scan(&divisionNumber1)
	fmt.Print("Enter First Division Number: ")
	fmt.Scan(&divisionNumber2)

	divisionOperation(divisionNumber1, divisionNumber2)
}