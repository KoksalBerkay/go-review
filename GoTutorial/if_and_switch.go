package main

import "log"

func main() {
	
	myNum := 100
	isTrue := false

	// If statement
	if myNum > 99 && !isTrue {
		log.Println("myNum is grater than 99 and isTrue set to false")
	} else if myNum < 100 && isTrue {
		log.Println("1")
	} else if myNum == 101 || isTrue{
		log.Println("2")
	} else if myNum > 1000 && isTrue == false{
		log.Println("3")
	}

	myVar := "cat"

	// Switch Statement
	switch myVar {
	case "cat":
		log.Println("myVar is set to cat")

	case "dog":
		log.Println("myVar is set to dog")
	
	case "fish":
		log.Println("myVar is set to fish")

	default:
		log.Println("myVar is something else")
	}
	
}