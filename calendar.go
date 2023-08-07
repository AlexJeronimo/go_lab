package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	//date := calendar.Date{}
	//date.Year = 2019
	//date.Month = 14
	//date.Day = 50
	//fmt.Println(date)
	/*
		err := date.SetYear(2019)
		if err != nil {
			log.Fatal(err)
		}

		err = date.SetMonth(5)
		if err != nil {
			log.Fatal(err)
		}

		err = date.SetDay(16)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(date)
		fmt.Println(date.Year())
		fmt.Println(date.Month())
		fmt.Println(date.Day())
	*/
	event := calendar.Event{}
	err := event.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(11)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(25)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event)
	fmt.Println(event.Year(), event.Month(), event.Day())

	err = event.SetTitle("Mom's birthday")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(event.Title())
}
