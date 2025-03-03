package main

import "fmt"

func FahrenheitToCelsius(fahrenheit float64) {
	var celsius int
	celsius = (int(fahrenheit) * 9 / 5) + 32
	fmt.Println("This is the conversion value from ", fahrenheit, "Celsius to Fahrenheit: ", celsius)
}

func main() {
	var fahrenheit float64
	fmt.Print("Enter Temperrature in Fahreheit: ")
	fmt.Scan(&fahrenheit)

	FahrenheitToCelsius(fahrenheit)
}
