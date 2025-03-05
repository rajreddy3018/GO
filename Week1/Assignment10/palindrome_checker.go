package main

import "fmt"

func palindromeChecker(number string) bool {

	var numToSlice = []rune(number)
	var lengthOfnumToSilce = len(numToSlice)
	
	for i,j:=0, (lengthOfnumToSilce - 1); i < j; i,j=i+1, j-1 {

		numToSlice[i],numToSlice[j] = numToSlice[j],numToSlice[i]
	}
	if (number) == string(numToSlice) {
		return true
	} else {
		return false
	}
}

func main() {
	
	var number string

	fmt.Print("Enter number to check if it is Palindrome or not:")
	fmt.Scan(&number)

	fmt.Printf("For the given number %s, Palindrome Checker is %t\n", number, palindromeChecker(number))


}