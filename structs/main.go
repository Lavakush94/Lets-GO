package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@outlook.com",
			zipCode: 98023,
		},
	}

	// fmt.Println(jim)
	jim.updateName("new name")
	jim.print()
}

func (ptp *person) updateName(nfn string) {
	(*ptp).firstName = nfn
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
