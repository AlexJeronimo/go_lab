package main

import (
	"fmt"
)

func main() {
	//err := errors.New("heigh can't be negative")
	//fmt.Println(err.Error())
	//fmt.Println(err)
	//.Fatal(err)

	height := -2.3333
	err := fmt.Errorf("a height of %0.2f is invalid", height)
	fmt.Println(err.Error())
	fmt.Println(err)
}
