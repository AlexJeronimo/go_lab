// 1
package main

import (
	"fmt"
)

func sayHi() {
	fmt.Println("Hi!")
}

var metersPerLiter float64

func paintNeeded(width float64, height float64) float64 {
	area := width * height
	//fmt.Printf("%.2f liters needed\n", area/10.0)
	return area / metersPerLiter
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
	metersPerLiter = 10.0
	fmt.Printf("%.2f", paintNeeded(4.2, 3.0))
}
