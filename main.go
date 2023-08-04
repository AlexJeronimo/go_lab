package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

func main() {
	/*
		var s magazine.Subscriber
		s.Rate = 4.99
		fmt.Println(s.Rate)

		subscriber := magazine.Subscriber{Name: "Aman Singh", Rate: 4.99, Active: true}
		fmt.Println("Name:", subscriber.Name)
		fmt.Println("Rate:", subscriber.Rate)
		fmt.Println("Active:", subscriber.Active)

		subscriber = magazine.Subscriber{Rate: 6.99}
		fmt.Println("Name:", subscriber.Name)
		fmt.Println("Rate:", subscriber.Rate)
		fmt.Println("Active:", subscriber.Active)

		var employee magazine.Employee
		employee.Name = "Joy Carr"
		employee.Salary = 6000
		fmt.Println(employee.Name)
		fmt.Println(employee.Salary)

		var address magazine.Address
		address.Street = "123 Oak St"
		address.City = "Omaha"
		address.State = "NE"
		address.PostalCode = "68111"
		fmt.Println(address)
	*/
	/*
		address := magazine.Address{Street: "123 Oak St", City: "Omaha", State: "NE", PostalCode: "68111"}
		subscriber := magazine.Subscriber{Name: "Aman Sigh"}
		subscriber.HomeAddress = address
		fmt.Println(subscriber.HomeAddress)
	*/
	/*
		subscriber := magazine.Subscriber{}
		fmt.Printf("%#v\n", subscriber.HomeAddress)

		subscriber.HomeAddress.PostalCode = "68111"
		fmt.Println("Postal Code:", subscriber.HomeAddress.PostalCode)
	*/
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.Address.Street = "123 Oak st"
	subscriber.Address.City = "Omaha"
	fmt.Println("Street:", subscriber.Address.Street)
	fmt.Println("City:", subscriber.Address.City)
	employee := magazine.Employee{Name: "Joy Carr"}
	employee.State = "OR"
	employee.PostalCode = "97222"
	fmt.Println("State:", employee.State)
	fmt.Println("Postal Code:", employee.PostalCode)
}
