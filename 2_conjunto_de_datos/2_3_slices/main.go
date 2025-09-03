package main

import "fmt"

func main() {

	//Los slices son apuntadores a un array, no poseen datos. Son porciones de arrays.
	//Array
	var beers [5]string = [5]string{"House lager", "House IPA", "House Cider", "KCBC", "Stout"}
	fmt.Println(beers)

	//slices
	houseBeers := beers[0:3]
	noHouseBeers := beers[3:5]
	houseBeers[1] = "House porter"

	fmt.Println(houseBeers)
	fmt.Println(noHouseBeers)

}
