package main

import (
	"encoding/json"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {

	myJson := `[
		{
			"first_name": "Clark",
			"last_name": "Kend",
			"hair_color": "Black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "wayne",
			"hair_color": "Brown",
			"has_dog": false
		}
	]`

	singleJson := `
	{
		"first_name": "Clark",
		"last_name": "Kend",
		"hair_color": "Black",
		"has_dog": true
	}`

	var p Person

	er := json.Unmarshal([]byte(singleJson), &p)

	log.Println(er)
	log.Printf("unmarshaled: %+v", p)

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("error while unmarshelling", err)
	}
	log.Printf("unmarshaled: %+v", unmarshalled)

}
