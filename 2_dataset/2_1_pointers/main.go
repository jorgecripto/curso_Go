package main

import "fmt"

func main() {

	//punteros: son variables que almacenan la dirección en memoria de un valor

	var name string = "Vortex INN"
	var punteroName *string = &name
	*punteroName = "Jorge Iván" // asignación de valor desde el puntero.

	fmt.Printf("El valor es: %v de tipo: %T guardado en: %v\n", name, name, &name)
	fmt.Printf("El valor es: %v de tipo: %T\n", punteroName, punteroName)

	//desreferenciación

	fmt.Printf("El valor es: %v de tipo: %T y la desrefereciación es: %s\n", punteroName, punteroName,
		*punteroName)

}
