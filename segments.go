package main

import (
	"fmt"
)

func main() {
	// var notes []string
	// notes = make([]string, 7)
	letters := []string{"a", "b", "c", "d"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	fmt.Println()
	for _, letter := range letters {
		fmt.Println(letter)
	}
}
