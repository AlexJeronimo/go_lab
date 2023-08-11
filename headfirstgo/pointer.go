package main

import (
	"fmt"
)

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(&myBoolPointer)
}

func main() {
	//amount := 6
	//fmt.Println(amount)
	//fmt.Println(&amount)
	/*
		var myInt int
		fmt.Println(reflect.TypeOf(&myInt))
		var myFloat float64
		fmt.Println(reflect.TypeOf(&myFloat))
		var myBool bool
		fmt.Println(reflect.TypeOf(&myBool))
	*/
	/*
		var myInt int
		myInt = 4
		var myIntPointer *int
		myIntPointer = &myInt
		fmt.Println(myIntPointer)
		fmt.Println(*myIntPointer)

		var myFloat float64
		myFloat = 98.6
		var myFloatPointer *float64
		myFloatPointer = &myFloat
		fmt.Println(myFloatPointer)
		fmt.Println(*myFloatPointer)

		var myBool bool
		myBool = true
		myBoolPointer := &myBool
		fmt.Println(myBoolPointer)
		fmt.Println(*myBoolPointer)
	*/
	/*
		myInt := 4
		fmt.Println(myInt)
		myIntPointer := &myInt
		*myIntPointer = 8
		fmt.Println(*myIntPointer)
		fmt.Println(myInt)
	*/

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	var myBool bool = true
	printPointer(&myBool)
}
