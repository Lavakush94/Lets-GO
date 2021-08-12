package main

import (
	"log"
	"sort"
)

func main() {
	myMap := make(map[string]string)
	myMap["dog"] = "samson"

	log.Println(myMap["none"])

	var myStrings []string

	names := []string{"one", "two", "three"}

	log.Println(names)

	myStrings = append(myStrings, "motu", "patalu", "john")

	sort.Strings(myStrings)

	log.Println(myStrings)
}
