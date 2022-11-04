package main

import (
	"fmt"
	"magazine"
)

func main() {
	var s magazine.Subscriber
	var e magazine.Employee
	s.Name = "김인하"
	s.Rate = 5.99
	s.Active = true
	e.Name = "김직원"
	// e.HomeAddress.City = "인천"
	e.City = "인천"
	fmt.Println(s.Name)
	fmt.Println(e.Name)
	fmt.Println(e.City)
}
