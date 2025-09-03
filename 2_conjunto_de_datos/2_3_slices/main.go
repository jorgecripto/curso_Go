package main

import "fmt"

func main() {

	//Los slices son apuntadores a un array, no poseen datos. Son porciones de arrays.
	//Array
	var beers [7]string = [7]string{"House lager", "House IPA", "House Cider", "House pilsner", "KCBC", "Baby Kittens", "Stout"}
	fmt.Println("Beers:", beers)

	//slices
	houseBeers := beers[0:4]
	noHouseBeers := beers[4:6]

	//Esta es una forma de cambiar el valor de una posición en particular.
	houseBeers[1] = "House porter"

	fmt.Println("Beers:", beers)
	fmt.Println("House Beers:", houseBeers)
	fmt.Println("No House Beers:", noHouseBeers)

	//Now we are going to see, len() and cap() funtions.
	//len(): it is the number of elements in my slice.
	//cap(): it is the number of elements from the original array, but starting from the position
	//when the slice was created.

	fmt.Println("Tamaño houseBeers:", len(houseBeers))
	fmt.Println("Capacidad noHouseBeers:", cap(noHouseBeers))

	fmt.Println("Tamaño noHouseBeers:", len(noHouseBeers))
	fmt.Println("Capacidad houseBeers:", cap(houseBeers))

	//Sin embargo, podemos modificar nuestros slices de acuerdo a su capacidad o simplemente creando
	//un nuevo array a partir del original.

	noHouseBeers = append(noHouseBeers, "Kona Big Wave")
	fmt.Println("New noHouseBeers on tap:", noHouseBeers)

	fmt.Println("Beers:", beers)

	//As we already know, slices are dynamic and they adapt according to our preferences.

	bottles := []string{"Corona", "Bud light", "Coors light", "Pumpkin night owl", "Budweiser"}
	fmt.Println(bottles)

	bottles = append(bottles, "Corona light", "Five Boroughs") //desde aca se le puede cambiar tamaño y capacidad.
	fmt.Println("Bottles updated:", bottles)

	cans := make([]string, 4, 5) //Se usa para optimizar código.
	cans[0] = "Five Boroughts"
	cans[1] = "High Noon Peach"
	cans[2] = "Athletic"
	cans[3] = "Guinness 0.0"
	fmt.Println(cans)

}
