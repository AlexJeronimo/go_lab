package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Actibe?", s.active)
}
func defaultSubscriber(name string) *subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)

	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)
}
