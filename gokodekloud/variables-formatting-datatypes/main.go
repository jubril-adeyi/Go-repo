package main

import (
	"fmt"
	"reflect"
)

func main() {
	var greeting string = "hello world"
	// /n means new line character and can prove useful in Print but Println fixes this
	fmt.Print(greeting, "\n")
	city := "london"
	girl := "lisa"
	//short variable decalration allows reasignments like :
	girl = "kemi"
	var age int = 8
	days := 9
	fmt.Println(girl, "is", age, "in", days, "days")
	var boy string = "harry"
	var grades int = 69
	// format specifiers
	// %v formats the value in default format meaning any datatypes
	// %d for int
	// %c for character
	// %s for strings
	// %t for boolean
	// %f for floats and
	// .2f for two decimal places floats
	// these to be used with printf from fmt

	fmt.Printf("nice to see you my boy, %v %v", boy, "\n")
	fmt.Printf("%v, you got %v in math, almost an A", boy, grades)
	fmt.Println("this is", city)

	// new internal block
	// use case:
	{
		girl = "amirah"
		country := "england"
		fmt.Printf("\nthis is a new block %v in %v \n", girl, country)
	}

	// declaring zero values, thr type needs to be declared first
	// use-case : checking if a varibale is empty or not initialized

	var i bool
	fmt.Println(i)
	// i is 0 if string its "", if float its 0.00 and if bool its false

	// finding types of variables datatype
	// %T format specifier
	grade := 60
	messsage := "passed"
	ischeck := true
	amount := 505.9

	fmt.Printf("variable grade = %v is of type %T \n", grade, grade)
	fmt.Printf("variable message = %v is of type %T \n", messsage, messsage)
	fmt.Printf("variable ischeck = %v is of type %T \n", ischeck, ischeck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)

	//reflect.typeof

	fmt.Println("using reflect.Typeof")

	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("hello"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(88.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))

	// with variables

	fmt.Printf("variable grade = %v is of type %v \n", grade, reflect.TypeOf(grade))
	fmt.Printf("variable grade = %v is of type %v \n", messsage, reflect.TypeOf(messsage))
	fmt.Printf("variable grade = %v is of type %v \n", ischeck, reflect.TypeOf(ischeck))
	fmt.Printf("variable grade = %v is of type %v \n", amount, reflect.TypeOf(amount))

}
