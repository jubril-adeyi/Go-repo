package main

import (
	"fmt"
	"strconv"
)

func main() {
	// converting variable to another datatypes
	// type casting = the value doesnt always remain the same after converting
	// int to float
	var e int = 90
	var f float64 = float64(e)
	fmt.Printf("%.2f\n", f)

	// so its basically just using the datatype name as a function and adding the original data in ()

	// float to integer

	var n float64 = 45.89
	var m int = int(n)
	fmt.Printf("%v\n", m)

	// strconv package  for conversion int to string using itoa

	var x int = 54
	var y string = strconv.Itoa(x)
	fmt.Printf("%q\n", y)

	// string conversion to int using atoi
	// var xx string = "400"
	var xx string = "400aa"
	yy, err := strconv.Atoi(xx)
	fmt.Printf("%v, %T \n", yy, yy)
	fmt.Printf("%v, %T", err, err)
}
