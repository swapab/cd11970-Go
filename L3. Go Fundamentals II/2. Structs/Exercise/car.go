package main

import (
	"fmt"
	"strconv"
)

type Car struct {
	make string
	year uint16
	used bool
}

func (c Car) describe() string {
	used := ""
	if c.used {
		used = "a used car"
	} else {
		used = "a new car"
	}
	return "This " + strconv.Itoa(int(c.year)) + " " + c.make + " is " + used
}

func car() {
	car1 := Car{make: "Audi", year: 2025, used: false}
	car2 := Car{make: "Tata", year: 1990, used: true}

	fmt.Println(car1)
	fmt.Println(car2)

	fmt.Println(car1.describe())
	fmt.Println(car2.describe())
}
