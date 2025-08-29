package main

import "fmt"

func main() {

	// ☑️ Declarar y asignar variables en Go en una misma linea.

	// var apple string = "Mi primera manzana en Go 🍎"
	// var banana string = "Mi primera banana en Go 🍌"
	// var orange string = "Mi primera naranja en Go 🍊"

	// ☑️ Declarar y asignar variables en bloque.

	// var (
	// 	cebolla string = "Mi primera cebolla en Go 🧅"
	// 	platano string = "Mi primer platáno en Go 🪴"
	// 	tomate  string = "Mi primer tomate en Go 🍅"
	// )

	// ☑️ Declarar y asignar variables en una sola linea.

	// var apple, banana, orange, cebolla, platano, tomate string = "🍎", "🍌", "🍊", "🧅", "🪴", "🍅"

	// ☑️ Declarar y asignar variables con operador de asignación de variable corta.

	apple, banana, orange, cebolla, platano, tomate := "🍎", "🍌", "🍊", "🧅", "🪴", "🍅"

	// fmt.Println(apple, "\n", banana, "\n", orange)
	// fmt.Println(cebolla, "\n", platano, "\n", tomate)

	fmt.Println(apple, banana, orange, cebolla, platano, tomate)

}
