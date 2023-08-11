// pass_fail informs is user complete examen
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// comment
	/*
		multiline comment
	*/
	fmt.Print("Entre a grade: ")
	reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failling"
	}
	fmt.Println("A grade of", grade, "is", status)
}
