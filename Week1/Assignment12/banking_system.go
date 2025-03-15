package main

import "fmt"

type BankAccount struct {
	AccountHolder string
	Balance       int
}

func main() {

	var person1 BankAccount
	var amount float64
	var action string

	fmt.Print("Enter Account Holder Name: ")
	fmt.Scan(&person1.AccountHolder)

	fmt.Print("Enter Account Holder Balance: ")
	fmt.Scan(&person1.Balance)

	fmt.Print("Amount: ")
	fmt.Scan(&amount)

	fmt.Print("Select Action (deposit/withdraw): ")
	fmt.Scan(&action)

	if action == "deposit" {
		Deposit(amount, person1)
	} else if action == "withdraw" {
		Withdraw(amount, person1)
	} else {
		fmt.Println("Please select valid action: deposit or withdraw")
	}
}

func Deposit(amount float64, accountHolderDetails BankAccount) {
	var Total float64
	Total = amount + float64(accountHolderDetails.Balance)

	fmt.Printf("Deposited %.2f into %s\n", Total, accountHolderDetails.AccountHolder)
}
func Withdraw(amount float64, accountHolderDetails BankAccount) {

	var Total float64
	if accountHolderDetails.Balance < 0 {
		fmt.Printf("Insufficient Funds in %s and amount available is %d", accountHolderDetails.AccountHolder, accountHolderDetails.Balance)
	} else {
		Total = float64(accountHolderDetails.Balance) - amount
		fmt.Printf("Remaining Balance %.2f for %s\n", Total, accountHolderDetails.AccountHolder)
	}

}
