package main

// import (
// 	"fmt"

// 	"rsc.io/quote"
// )

// func main() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println(quote.Go())
// }

import (
	"fmt"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}
