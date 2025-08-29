package main

import "fmt"

func main() {

	//Los datos tipo 1️⃣bool pueden ser verdaderos (true) o falsos (false).
	var exit bool = true
	/* Vamos a usar este nuevo print de formato para imprimir en pantalla usando verbos como
	indicadores de posición dandole los valores para que me retornen un mensaje tipo string*/
	fmt.Printf("El valor es: %v y el tipo es: %T\n", exit, exit)

	//Los datos de tipo 2️⃣string siempre deben estar entre comillas dobles.

	var name string = "Vortex Innovation is the future"
	fmt.Printf("El valor es: %v y el tipo es: %T\n", name, name)

	//Los datos de tipo 3️⃣numeric son mas variados y los vamos a enlistar abajo
	//en un comentario de bloque.

	/* uint

	👽 uint8 unsigned 8-bit integers (0 to 255)
	👽 uint16 unsigned 16-bit integers (0 to 65535)
	👽 uint32 unsigned 32-bit integers (0 to 4294967295)
	👽 uint64 unsigned 64-bit integers (0 to 18446744073709551615)

	int

	👽 int8 signed 8-bit integers (-128 to 127)
	👽 int16 signed 16-bit integers (-32768 to 32767)
	👽 int32 signed 32-bit integers (-2147483648 to 2147483647)
	👽 int64 signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	byte // alias for uint8

	rune // alias for int32 // represents a Unicode code point

	float32 float64
	*/

	var age uint8 = 38
	fmt.Printf("El valor es: %v y el tipo es: %T\n", age, age)

	var kode rune = 'a'
	fmt.Printf("El valor es: %v y el tipo es: %T\n", kode, kode)

	//Tener cuidado de hacer operaciones con diferentes tipos de datos ya que el compilador
	//no lo dejara, a menos que hagas casting lo que te permitira hacer la operación pero jamas
	//cambiar el tipo de dato de la variable. Esto es el mejor ejemplo de que Go es estaticamente tipado

	var x uint8 = 25
	var y uint32 = 1265
	z := uint32(x) * y
	fmt.Printf("El valor es: %v y el tipo es: %T\n", z, z)

	/*Algo sumamente util cuando asignamos algún tipo de lógica pero no la queremos declarar pero
	la queremos mantener util para ser llamada en cualquier momento y que no afecte nuestro
	código es usar el operador blank que es raya al piso*/

	_ = uint32(x) * y

}
