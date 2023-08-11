package main

import (
	"fmt"
)

func manyReterns() (int, bool, string) {
	return 1, true, "hello"
}

func main() {
	myInt, myBool, myString := manyReterns()
	fmt.Println(myInt, myBool, myString)
}
