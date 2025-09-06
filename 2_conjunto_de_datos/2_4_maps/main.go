package main

import "fmt"

func main() {

	//we can use the reserved funtion make to build a map
	things := make(map[string]string)
	things["Book"] = "📖"
	things["Pencil"] = "✏️"
	things["Pen"] = "🖊️"

	//fmt.Println(things)

	//this is another way to build it.
	office := map[string]string{
		"desktop":    "🖥️",
		"paperclips": "🖇️",
		"folder":     "📂",
		"worker":     "🧑‍🏭",
	}

	//fmt.Println(office)

	//we can use the reserved word delete to eliminate one of the elements from our map.
	delete(office, "worker")
	//fmt.Println(office)

	//we can find out the content in the map using this code:
	//we use the word "ok" to receive a bool.

	content, ok := office["worker"]
	fmt.Println(content, ok)

	cont, ok := office["folder"]
	fmt.Println(cont, ok)
}
