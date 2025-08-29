package main

import "fmt"

func main() {

	//operadores de comparación >, <, ==, !=, >=, <=
	//Estos me permiten comparar y me darán una respuesta booleana.

	// fmt.Println(6 > 5)
	// fmt.Println(10 < 9)
	// fmt.Println(6 == 66)
	// fmt.Println(6 != 5)
	// fmt.Println(60 >= 61)
	// fmt.Println(15 <= 50)

	//operadores lógicos &&, ||

	fmt.Println("hello, welcome to the the Bar, could you type your age?")
	var age uint8
	fmt.Scan(&age)

	//Primer bloque de código para el &&, para acticarlo descomente.

	// if age >= 18 && age <= 60 {
	// 	fmt.Println("Muchas Gracias, puedes ingresar")
	// } else if age < 18 {
	// 	fmt.Println("Lo siento, NO puedes ingresar porque eres menor de edad")
	// } else if age > 60 {
	// 	fmt.Println("Lo siento No puedes ingresar porque hay un evento privado para jovenes")
	// }

	//Segundo bloque de c+odigo para el ||

	if age == 18 || age > 18 && age <= 60 {
		fmt.Println("Muchas Gracias, puedes ingresar")
	} else if age < 18 || age > 60 {
		fmt.Println("Lo siento, NO puedes ingresar")
	}

	//operador lógico Unario: !, este es el operador que usamos para negar

	fmt.Println(!(4 == 4))
	fmt.Println(!(4 != 5))

}
