package main

import "fmt"

func Switch(value int) string {
	switch value {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	default:
		return "Default"
	}
}

func IfElse(num int) string {

	Default := "Integer value not provided"

	if num == 1 {
		return "one"
	} else if num == 2 {
		return "two"
	} else if num == 3 {
		return "three"
	} else if num == 4 {
		return "four"
	} else {
		return Default
	}
}

func main() {
	fmt.Println(Switch(3))
	fmt.Println(IfElse(4))
	fmt.Println(IfElse('t'))
}
