package main

import (
	"log"
	"time"
)

var s = "seven"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	user := User{
		FirstName: "Salmon",
		LastName:  "Khan",
	}

	var myVar myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{
		FirstName: "Motu",
	}

	log.Println("myvar is set to", myVar2.printFirstName())

	log.Println(user.FirstName)
}
