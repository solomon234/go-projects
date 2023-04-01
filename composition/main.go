package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Num of Passengars:", v.NumberOfPassengers)
	fmt.Println("Num of wheels:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year", c.Year)
	fmt.Println("IsElectric", c.IsElectric)
	fmt.Println("IsHybrid", c.IsHybrid)
}

func main() {
	suv := Vehicle{
		NumberOfWheels:     4,
		NumberOfPassengers: 6,
	}

	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90 T8",
		Year:       2020,
		IsElectric: false,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	volvoXC90.show()
	fmt.Println()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "Model X",
		Year:       2022,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	teslaModelX.show()
	fmt.Println()

}
