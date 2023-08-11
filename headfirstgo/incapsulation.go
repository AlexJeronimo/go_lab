package main

import (
	"errors"
	"fmt"
	"log"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetYear(year int) error {
	if year < 1 {
		return errors.New("Invalid year")
	}
	d.Year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid day")
	}
	d.Day = day
	return nil
}

func main() {
	//date := Date{Year: 2019, Month: 5, Day: 27}
	//fmt.Println(date)
	date := Date{}
	//err := date.SetYear(0)
	//err := date.SetMonth(15)
	err := date.SetDay(14)
	if err != nil {
		log.Fatal(err)
	}
	//date.SetYear(2019)
	//fmt.Println(date.Year)
	//fmt.Println(date.Month)
	fmt.Println(date.Day)
}
