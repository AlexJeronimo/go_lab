package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Milliliters) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}

func (g Gallons) ToLiters() Liters {
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters {
	return Milliliters(g * 3785.41)
}

func main() {
	/*
		var carFuel Gallons
		var busFuel Liters
		carFuel = Gallons(10.0)
		busFuel = Liters(240.0)
		fmt.Println(carFuel, busFuel)
		carFuel = Gallons(Liters(40) * 0.264)
		busFuel = Liters(Gallons(63) * 3.785)
		fmt.Printf("Galons: %0.1f Liters: %0.1f\n", carFuel, busFuel)
	*/
	/*
		carFuel := Gallons(1.2)
		busFuel := Liters(4.5)
		carFuel += ToGallons(Liters(40.0))
		busFuel += ToLiters(Gallons(30.0))
		fmt.Printf("Car Fuel: %0.1f gallons\n", carFuel)
		fmt.Printf("Bus Fuel: %0.1f liters\n", busFuel)
	*/
	soda := Liters(2)
	fmt.Printf("%0.0f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)
	fmt.Printf("%0.0f milliliters equals %0.3f gallons\n", water, water.ToGallons())
}
