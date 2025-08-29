package main

import "fmt"

// Las constantes son mejor trabajarlas a nivel de package para que siempre esten disponibles,
// en todas mis funciones. Si es de manera grupal mucho mejor.

const (
	celestialBodyFactor = 0.182
	planetsNumber       = 10
)

// Le podemos incrementar su valor de asignación a un bloque de constantes iota.

const (
	january = iota + 1
	february
	march
	april
	may
	june
	july
	august
	september
	october
	november
	december
)

func main() {

	const hearthGravity float64 = 9.8
	fmt.Println(hearthGravity, "m/s^2")

	// Varias constantes también se puede declarar y asignar en una misma linea.
	// El operador de asignación de variable corta no se puede usar en las constantes,
	// pero si se puede inferir el tipo de dato manteniendo la palabra reservado const.

	const factor = 0.165
	const moonGravity = (factor * hearthGravity)
	fmt.Println(moonGravity, "m/s^2")

	fmt.Println(celestialBodyFactor, planetsNumber)

	fmt.Println(november)

}
