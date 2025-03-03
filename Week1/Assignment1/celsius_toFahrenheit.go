package celsiutoFahren

import "fmt"

func CelsiusToFahrenheit(celsius int) {
	
	var fahrenheit float64
	fahrenheit=(float64(celsius) * 9/5) + 32
	fmt.Println("This is the conversion value from ",celsius, "Celsius to Fahrenheit: ",fahrenheit)

}

func main() {
	var celsius int
	fmt.Print("Enter Celsiu:")
	fmt.Scan(&celsius)
	CelsiusToFahrenheit(celsius)
}
