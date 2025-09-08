package main

import "fmt"

type Person struct {
	Name     string
	LastName string
	Age      uint8
	Weight   uint8
}

func main() {

	CEO := Person{
		Name:     "Andres",
		LastName: "Ardila",
		Age:      34,
		Weight:   95,
	}

	// CTO := Person{
	// 	Name:     "Jorge Iván",
	// 	LastName: "Guevara",
	// 	Age:      38,
	// 	Weight:   83,
	// }

	// caveman := Person{"Jesus", "Kiyosaky", 25, 70}

	// fmt.Println(CEO)
	// fmt.Println(CTO)
	// fmt.Println("El nombre del Cavernicola es:", caveman)

	//we can print it this way to have tha name of the elements and the value.

	//fmt.Printf("%+v", CEO)

	//with structs we can work with pointers as well.

	boss := &CEO
	fmt.Println(boss)

	//we can change certains values this way.
	(*boss).Age = 35 //also = boss.Age = 40.
	fmt.Println(boss)
}
