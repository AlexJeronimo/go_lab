package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func main() {
	/*
		var myStruct struct {
			number float64
			word   string
			toggle bool
		}
		fmt.Printf("%#v\n", myStruct)
		myStruct.number = 3.14
		myStruct.word = "pie"
		myStruct.toggle = true
		fmt.Println(myStruct.number)
		fmt.Println(myStruct.word)
		fmt.Println(myStruct.toggle)
	*/
	/*
		var subscriber struct {
			name   string
			rate   float64
			active bool
		}
		subscriber.name = "Aman Singh"
		subscriber.rate = 4.99
		subscriber.active = true
		fmt.Println("Name:", subscriber.name)
		fmt.Println("Monthly rate:", subscriber.rate)
		fmt.Println("Active?", subscriber.active)
	*/
	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 subscriber
	subscriber2.name = "Beth Ryan"
	fmt.Println("Name:", subscriber2.name)
}
