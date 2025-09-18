package main

import (
	"errors"
	"fmt"
	"strconv"
)

var errNotFound = errors.New("what you are searching, does not exist")

var itemsOffice = map[int]string{
		1:"🖥️",
		2:"🖇️",
		3:"📂",
		4:"🧑‍🏭",
	}

func main() {
	found, err := search("5")
	if errors.Is(err, errNotFound){
		fmt.Println("We could handle the error, this is the real error:", err)
		return
	}
	if err != nil {
		fmt.Println("search()", err)
	}
	fmt.Println(found)
}

func search(key string) (string, error)  {
	num, err := strconv.Atoi(key)
	if err != nil {
		return "", fmt.Errorf("strconv.Atoi(): %w", err)
	}	

	item, err := findItem(num)
	if err != nil {
		return "", fmt.Errorf("findItem(): %w", err)
	}
	return item, nil
}

func findItem(id int) (string, error)  {

	value, ok := itemsOffice[id]
	if !ok {
		return "", errNotFound
	}
	return value, nil
	
}