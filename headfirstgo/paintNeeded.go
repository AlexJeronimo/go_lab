// 1
package main

import (
	"fmt"
	"log"
)

/*
	func sayHi() {
		fmt.Println("Hi!")
	}
*/
var metersPerLiter float64

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	//fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / metersPerLiter, nil
}

func main() {
	//var width, height, area float64
	//width = 4.2
	//height = 3.0
	//area = width * height
	//fmt.Println(area/10.0, "liters needed")
	//fmt.Printf("%.2f liters needed\n", area/10.0)
	//sayHi()

	//paintNeeded(4.2, 3.0)
	//paintNeeded(5.2, 3.5)
	metersPerLiter = 10.

	var amount, total float64
	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		//fmt.Println(err)
		log.Fatal(err)
	} //else {
	//	fmt.Printf("%0.2f liters needed\n", amount)
	//}
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	//amount, err = paintNeeded(5.2, 3.5)
	//fmt.Printf("%0.2f liters needed\n", amount)
	//total += amount
	//fmt.Printf("Total: %0.2f liters\n", total)

	//fmt.Printf("%.2f", paintNeeded(4.2, 3.0))
}
