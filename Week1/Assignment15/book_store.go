package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
}

func main() {

	var book1 Book

	fmt.Print("Enter Book Title: ")
	fmt.Scan(&book1.Title)
	fmt.Print("Enter Book Author: ")
	fmt.Scan(&book1.Author)
	fmt.Print("Enter Book Price: ")
	fmt.Scan(&book1.Price)

	if book1.Price > 50 {
		aboveFiftyDollars(book1)
	} else {
		fmt.Println("No discount! Since the book price is less than 50 dollars")
	}

}

func aboveFiftyDollars(price Book) {
	price.Price *= 0.5
	fmt.Printf("Discount is Applied! Pay this amount: %.2f\n", price.Price)
}
