package main

import "fmt"

const colombianLegalAge uint8 = 18
const usLegalAge uint8 = 21

func main() {

	//You can define the variable first and then do the if flow control.
	
	var age uint8 = 21

	if age < usLegalAge {
		fmt.Println("You are not of age yet")
	} else if age >= usLegalAge {
		fmt.Println("Your are of legal age, already")
	}

	//Also, you can define the var inside the flow control.

	if animal := "🐱"; animal == "🐕"  {
		fmt.Printf("It is a %s\n", animal)		
	} else if animal == "🐱" {
		fmt.Printf("It is a %s\n", animal)
	} else {
		fmt.Println("This is not a domestic animal")
	}

	//switch in a boolean case.
	
	var colombianAge uint8 = 19
	switch {
	case colombianAge < colombianLegalAge:
		fmt.Println("You are not of age in Colombia")
	default:
		fmt.Println("You are of age in Colombia")
	}

	animal2 := "👩‍🎨"

	switch animal2 {
	case "🐦":
		fmt.Printf("This %s can fly\n", animal2)
	case "🐶":
		fmt.Printf("This %s can run\n", animal2)
	case "🐟":
		fmt.Printf("This %s can swim\n", animal2)
	default:
		fmt.Printf("This %s is not an animal\n", animal2)
		
	}

	var usAge uint8 = 20
	usBar := false
	colombianBar := false

	switch {
	case colombianAge < colombianLegalAge && colombianBar:
		fmt.Println("You are not allowed to come in to the Colombian Bar")
	case colombianAge > colombianLegalAge && colombianBar:
		fmt.Println("You are allowed to come in to the Colombian Bar")
	case usAge < usLegalAge && usBar:
		fmt.Println("You are not allowed to come in to the American bar")
	case usAge > usLegalAge && usBar:
		fmt.Println("You are allowed to come in to the American bar")
	default:
		fmt.Println("This is neither Colombia nor US")
	}

}