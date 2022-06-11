package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// Map
	myMap := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)

	// Slice
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	log.Println(numbers[0:2])

	names := []string{"one", "seven", "zero"}
	log.Println(names)

}
