package main

import "fmt"

/*barra inclinada + asterisco se usa para comentarios en bloque cuando no queremos que una
parte de nuestro código este disponible como el ejemplo de las constantes de meses abajo*/

/*const (
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
)*/

func main() {
	// las barras inclinadas dobles se usan para comentarios de una sola linea.

	// podemos comentar declaraciones como la siguiente:
	//name es el nombre del dueño de este repositorio.
	var name string = "Jorge Iván"
	fmt.Println(name)
}
