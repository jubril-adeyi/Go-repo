package main

import "fmt"

//constant can be declared globally outside the main func

const pi float64 = 3.14

func main() {
	// declaring constants using comnst <const name> <datatype>
	// but the datatype can be skipped its is automatically inferred also
	const boy string = "male"
	const girl = "female"
	const surname = "adeyi"
	const iswhite = false

	fmt.Print(boy)
	fmt.Print("\n", girl, "\n")
	fmt.Printf("%v: %T \n", surname, surname)
	fmt.Printf("%v: %T \n", iswhite, iswhite)

	// const value cannot be reassigned
	// also cannot declare constant and not initialize with no value like in var
	// short hand := doesnt work for const

	// use case

	radius := 10.00
	var area float64

	area = pi * radius * radius
	fmt.Print(area)

	//use of operators (*) and operands are the variables nad values

}
