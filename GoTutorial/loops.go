package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Joe", "Mama", "joe@mama.com", 20})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 45})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}