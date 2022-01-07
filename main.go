package main

import "fmt"

type Vehicle struct {
	NumbersOfWheels    int
	NUmberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	isElectric bool
	isHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Numbers of passengers:", v.NUmberOfPassengers)
	fmt.Println("Numbers of wheels:", v.NumbersOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:",c.Make)
	fmt.Println("Model:",c.Model)
	fmt.Println("Year:",c.Year)
	fmt.Println("Is Electric:",c.isElectric)
	fmt.Println("Is Hybrid:",c.isHybrid)
	c.Vehicle.showDetails()
}
func main() {
	suv:=Vehicle{
		NumbersOfWheels: 4,
		NUmberOfPassengers: 6,
	}

	volvoXC90:=Car{
		Make:"Volvo",
		Model:"XC90 T8",
		Year: 2021,
		isElectric: false,
		isHybrid: true,
		Vehicle: suv,
	}

	volvoXC90.show()
}