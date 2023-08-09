package main

import (
	"fmt"

	"github.com/headfirstgo/prose"
)

func main() {
	//fmt.Println(strings.Join([]string{"05", "14", "2018"}, "/"))
	//fmt.Println(strings.Join([]string{"state", "of", "the", "art"}, "-"))
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photot of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photot of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents"}
	fmt.Println("A photot of", prose.JoinWithCommas(phrases))
}
