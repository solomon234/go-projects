package main

import (
	"log"
	"myapp/staff"
)

var employees = []staff.Employee{
	{FirstName: "John", LastName: "Smith", Salary: 30000, FullTime: true},
	{FirstName: "Saly", LastName: "Jones", Salary: 50000, FullTime: true},
	{FirstName: "MArk", LastName: "Smithers", Salary: 60000, FullTime: true},
	{FirstName: "Allan", LastName: "Anderson", Salary: 15000, FullTime: true},
	{FirstName: "Margaret", LastName: "Carter", Salary: 5000, FullTime: true},
}

func main() {
	myStaff := staff.Office{
		AllStaff: employees,
	}

	//log.Println(myStaff.All())
	log.Println("Overpaid staff", myStaff.Overpaid())
	log.Println("underpaid staff", myStaff.Underpaid())

}
