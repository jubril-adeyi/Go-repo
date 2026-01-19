package main

import "fmt"

// if statements, switch-case and fallthrough, for loops and break and continue statement

func main() {
	// fmt.Println("control flow")

	age := 17

	// if-else, else if and else combo here

	// if age >= 18 {
	// 	fmt.Println("allow, adult")
	// } else if age == 17 {
	// 	fmt.Println("close but denied jor!")
	// } else {
	// 	fmt.Println("Below 18! access denied")
	// }

	// using switch-case staetemnt
	fmt.Println("using switch case for control flow \n")
	// age = 18
	switch {
	case age >= 18:
		fmt.Println("allow, adult")
	case age == 17:
		fmt.Println("close but denied jor!")
	default:
		fmt.Println("below 18! access denied jor")
	}

	// var age2 int = 21
	// switch age2 {
	// case 15:
	// 	fmt.Println("you cannot have alcohol, you are not  old enough")
	// case 22:
	// 	fmt.Println("you can have alcohol, you are old enough")
	// default:
	// 	fmt.Println("just dont drink, i'm not sure")
	// }

	var i int = 800
	switch i {
	case 10:
		fmt.Println("i is 10")
	// case takeing 2 args
	case 100, 200:
		fmt.Println("i is either 100 or 200")
	default:
		fmt.Println("i is neither 0, 100, or 200")
	}

	// fallthrough keyord in switch cases
	// to force the execution to fall through all the cases that have fallthrough within them.

	i = 10
	switch i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10")
		fallthrough
	case 20:
		fmt.Println("20")
		fallthrough
	default:
		fmt.Println("default")
	}

	// conditions within the case block
	var a, b int = 10, 20
	switch {
	case a+b == 30:
		fmt.Println("equal to 30 ")
	case a+b <= 30:
		fmt.Println("lessthan or equal to 30")
	default:
		fmt.Println("Greater than 30")
	}

	// Loops
	// for loops
	// for initializaton(assignent or declaration of value); condition using conditionals; post(action){
	// 	statement
	// }
	var j int
	// for j = 1; j <= 3; j++ {
	// 	// value is 1, the range is within 3 and the action is the increment by 1
	// 	fmt.Println("hello world")
	// }

	j = 1
	for j <= 3 {
		fmt.Println(j * 8)
		j += 1

	}

	// infinite loop
	// for {
	// 	j++
	// }
	// fmt.Println(j)

	// using break statement to exit for loop used here with an if statement
	for j = 2; j <= 10; j++ {
		if j == 8 {
			break
		}
		fmt.Println(j)
	}

	// continue statement used to skip when the if statement condition is met and then keep going
	for j = 5; j <= 15; j++ {
		// so here it skips 12 an moves on to 13..
		if j == 12 {
			continue
		}
		fmt.Println(j)
	}
}
