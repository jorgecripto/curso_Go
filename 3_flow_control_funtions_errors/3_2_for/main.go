package main

import "fmt"

func main() {

	//This is the regular for (legacy statement).
	//i := 0 It is the start statement that is evaluated before each iteration.
	//i < count es la segunda condición que es evaluada antes de cada iteración
	//y finalmente tenemos la declaración posterior i++ (i = i + 1) que es ejecutada cada vez que finaliza
	//una iteración.
	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	bottles := []string{"Corona", "Bud light", "Coors light", "Pumpkin night owl", "Budweiser"}
	for i, can := range bottles {
		fmt.Println("Index:", i+1, "Name:", can)		
	}

	office := map[string]string{
		"desktop":    "🖥️",
		"paperclips": "🖇️",
		"folder":     "📂", 
		"worker":     "🧑‍🏭",
	}

	for key, value := range office {
		fmt.Println("Index:", key, "Name:", value)
	}
}