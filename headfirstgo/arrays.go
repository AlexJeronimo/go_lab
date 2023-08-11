package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	/*
		index := 1
		fmt.Println(index, notes[index])
		index = 3
		fmt.Println(index, notes[index])

		for i := 0; i <= 2; i++ {
			fmt.Println(i, notes[i])
		}

		fmt.Println(len(notes))
		primes := [5]int{2, 3, 5, 7, 11}
		fmt.Println(len(primes))
	*/
	/*
		for i := 0; i < len(notes); i++ {
			fmt.Println(i, notes[i])
		}
	*/

	for index, note := range notes {
		fmt.Println(index, note)
	}
}
