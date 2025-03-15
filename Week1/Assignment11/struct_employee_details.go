package main

import "fmt"

func displayEmployeeDetails(Details EmployeeDetails) {
	fmt.Println("\nEmployee Details are Below:")
	fmt.Println("Employee Name: ", Details.Name)
	fmt.Println("Employee Age: ", Details.Age)
	fmt.Println("Employee Gender: ", Details.Gender)
	fmt.Println("Employee Salary: ", Details.Salary)
	fmt.Println("Employee PhoneNumber: ", Details.PhoneNumber)
}

type EmployeeDetails struct {
	Name        string
	Age         int
	Gender      string
	Salary      int
	PhoneNumber int
}

func main() {
	var employee1, employee2 EmployeeDetails
	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&employee1.Name)

	fmt.Print("Enter Employee Age: ")
	fmt.Scan(&employee1.Age)

	fmt.Print("Enter Employee Gender: ")
	fmt.Scan(&employee1.Gender)

	fmt.Print("Enter Employee Salary: ")
	fmt.Scan(&employee1.Salary)

	fmt.Print("Enter Employee PhoneNumber: ")
	fmt.Scan(&employee1.PhoneNumber)

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&employee2.Name)

	fmt.Print("Enter Employee Age: ")
	fmt.Scan(&employee2.Age)

	fmt.Print("Enter Employee Gender: ")
	fmt.Scan(&employee2.Gender)

	fmt.Print("Enter Employee Salary: ")
	fmt.Scan(&employee2.Salary)

	fmt.Print("Enter Employee PhoneNumber: ")
	fmt.Scan(&employee2.PhoneNumber)

	displayEmployeeDetails(employee1)
	displayEmployeeDetails(employee2)
}
