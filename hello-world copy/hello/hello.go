package main

import (
	"fmt"
	greeting "greetings/greet"
	"log"
)

func main() {
	// name := "jubril"
	// surname := "adeyi"
	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())
	// fmt.Println(name, surname)

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greeting.Greet("")
	fmt.Print(message)

	if err != nil {
		log.Fatal(err)
	}
}

// message := greeting.Greets(name)

// message2 := cursing.Curse(name)
// fmt.Println(message)
// fmt.Println(message2)
