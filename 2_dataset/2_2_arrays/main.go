package main

import "fmt"

func main() {

	//el 3 indica el tamaño del array

	var colors [3]string = [3]string{"💛", "💙", "🔴"}
	fmt.Println(colors)

	//Esta es otra forma de declarar los arrays

	nameBa := [3]string{"Pedro", "Edagar", "Fred"}
	fmt.Println(nameBa)

	//Podemos declarar y asignar por separado nuestro array

	var sueldo [3]float64
	sueldo[0] = 854.12
	sueldo[1] = 954.36
	sueldo[2] = 1200.00
	fmt.Println(sueldo)

	//Inferencia del tamaño del array, sin embargo, sigue siendo de tamaño fijo.

	nameBa2 := [...]string{"Pedro", "Edagar", "Fred", "Gio", "Aoife"}
	fmt.Println(nameBa2)

}
