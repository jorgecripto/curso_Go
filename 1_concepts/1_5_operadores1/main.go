package main

import "fmt"

func main() {
	//operadores aritmético (), *, /, %, +, -

	var x = 2 + 3*2 //tener en cuenta la jerarquia de las operaciones.
	fmt.Println(x)

	//operadores de asignación =, +=, -=, *=, /=, %=

	var g int = 25
	g += 5
	fmt.Println(g)

	//declración post-incremento y post-decremento: ++, --
	//(no son una expresión sino una declaración)

	var h int = 10
	h++
	fmt.Println(h)

}
