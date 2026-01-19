package main

import "fmt"

func main() {
	// fmt.Println("control flow")

	age := 17

	// if-else, else if and else combo here

	if age >= 18 {
		fmt.Println("allow, adult")
	} else if age == 17 {
		fmt.Println("close but denied jor!")
	} else {
		fmt.Println("Below 18! access denied")
	}
}
