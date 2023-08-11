package main

import "fmt"

type Number int
type MyType string

func (n *Number) Double() {
	*n *= 2
}

func (m MyType) method() {
	fmt.Println("Method with value receiver")
}

func (m *MyType) pointerMethod() {
	fmt.Println("Method with pointer reciever")
}

func main() {
	number := Number(4)
	fmt.Println("Original velue of number:", number)
	number.Double()
	fmt.Println("number after calling Double:", number)

	fmt.Println()

	value := MyType("a value")
	pointer := &value
	value.method()
	value.pointerMethod()
	pointer.method()
	pointer.pointerMethod()
	fmt.Println(*pointer)
	fmt.Println(&value)
}
