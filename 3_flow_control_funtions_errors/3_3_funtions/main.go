package main

import "fmt"

func main() {

	values := []int{4, 10, 22, 54, 85, 71, 101, 21}

	//You have to remember that although the function, you are calling in the main, is proccesing
	//what it is created for, you must save the return in a variable and then print it in console.
	result := sum(5, 6)
	fmt.Println(result)

	// This is an option to return what you want to filter without creating another function.

	// newValuesResult := filter(values, func(value int)bool{return value > 22})
	// fmt.Println(newValuesResult)

	// This is a neater code but you have to create the functions first and then do the callback.

	newValuesResult := filter(values, lessThanEighty)
	fmt.Println(newValuesResult)

	// here, we call the function plus.

	fmt.Println(plus(2, 5, 9))
	fmt.Println(plus(252, 5879, 25, 598, 2635))

	// Anonymous function.

	greet := func(){
	fmt.Println("✋ Hello my love, I am learning Go")
	}

	greet()

}

func sum(a, b int) int {
	return a + b
}

//Functions are data types as well.

func filter(values []int, callback func(int) bool ) []int  {
	var resultFilter []int
	// resultFilter := make([]int, 0, len(values)) 

	for _, value := range values {
		if callback(value){
			resultFilter = append(resultFilter, value)
		}
		
	}
	return resultFilter
}

// func greatherToTwentytwo(value int) bool  {

// 	return value > 22
	
// }

func lessThanEighty(value int) bool  {

	return value < 80
	
}

// Variatic and anonymous functions: Let´s go back to our function sum but in this case we want
// to have more numbers within the function.

func plus(numbers ...int) int {  //we call this function in the function main
	var total int

	for _, number := range numbers {
		total += number
		
	}
	return total
}

// The anonymous function will be called within of the main. So, Let´s scroll up.