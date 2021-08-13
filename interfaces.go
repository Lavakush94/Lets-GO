package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Dog) Says() string {
	return "Woof"
}

type Grorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func (d *Grorilla) NumberOfLegs() int {
	return 2
}

func (d *Grorilla) Says() string {
	return "Ugh"
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)

	gorilla := Grorilla{
		Name:          "Gorilla",
		Color:         "Brown",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), " and has ", a.NumberOfLegs(), "legs")
}
