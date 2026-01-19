package main

import "fmt"

func main() {
	// var name string
	// var surname string
	// fmt.Print("Enter you name here: ")
	//taking user input
	// it takes the format string and then object arguments
	// the string contains format specifier that the function used to format the output string

	// here the format specifier for a strimng is used (%s), then & used before the variablke
	// this is how the Scanf function from fmt works in gp
	// fmt.Scanf("%s", &name)
	// fmt.Println("name: ", name)

	//multiple inputs the input picks up sequentually and picked in orde rspecified in code

	// fmt.Scanf("%s %s", &name, &surname)
	// fmt.Println(name, surname)

	// scanf returns count(nunber of arguments that the function writes to)
	// and err values (error thrown duringtheexecution of the function)

	var a string
	var b int

	fmt.Print("Enter a string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)

	fmt.Println("count : ", count)
	fmt.Println("error: ", err)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

}
