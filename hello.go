package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var whatToSay string
	whatToSay = "good morning"

	fmt.Println(whatToSay)

	var i int
	i = 8

	fmt.Println(i)

	whatWasSaid1, whatWasSaid2 := saySomething()
	fmt.Println(whatWasSaid1, whatWasSaid2)

	callByReference(&whatToSay)
	fmt.Println(whatToSay)
}

func saySomething() (string, string) {
	return "something", "Went wrong"
}

func callByReference(s *string) {
	p := "runs fine"
	*s = p
}
