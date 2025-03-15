package main

import "fmt"

type Vehicle interface {
	speed() int;
}

type Car struct {
	Make string;
	Model string;
	MaxSpeed int;
}

type Bike struct {
	Brand string;
	Type string;
	MaxSpeed int;
}

func (c Car) speed() int {
	return c.MaxSpeed
}

func (b Bike) speed() int {
	return b.MaxSpeed
}

func main() {

	car1 := Car{Make: "Suzuki", Model: "Maruthi", MaxSpeed: 200  } 
	bike1 := Bike{Brand: "Tata", Type: "Spendor", MaxSpeed: 100  } 

	var vehicle Vehicle

	vehicle = car1

	fmt.Printf("Speed of car %dkm/h", vehicle.speed())

	vehicle = bike1
	fmt.Printf("Speed of bike %dkm/h", vehicle.speed())

}