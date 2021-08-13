package main

import "log"

func main() {
	num := 5
	if num%2 == 0 {
		log.Println("even number")
	} else {
		log.Println("odd number")
	}

	// switch
	caseVariable := "cat"

	switch caseVariable {
	case "cat":
		log.Println("this is cat")
	case "dog":
		log.Println("this is dog")
	case "fish":
		log.Println("this is fish")
	default:
		log.Println("cat is something else")
	}

	//loop

	type User struct {
		FristName string
		LastName  string
		Email     string
		Age       int
	}

	for i := 1; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "cat", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	var users []User

	users = append(users, User{"John", "the don", "john@gmail.com", 25})
	users = append(users, User{"Motu", "", "motu@gmail.com", 125})
	users = append(users, User{"Patlu", "", "patalu@gmail.com", 125})
	users = append(users, User{"jhatka", "dr", "jhatka@gmail.com", 25})
	users = append(users, User{"Ghaseeta", "Ram", "ram@gmail.com", 25})

	for _, l := range users {
		log.Println(l.FristName)
	}
}
