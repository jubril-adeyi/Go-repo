package main

import (
	"fmt"
	cursing "greetings/curse"
	greeting "greetings/greet"
	"log"

	"rsc.io/quote"
)

func main() {
	name := "jubril"
	surname := "adeyi"
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(name, surname)

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greeting.Greet("")
	fmt.Print(message)

	if err != nil {
		log.Fatal(err)
	}
	message1 := greeting.Greets(name)

	message2 := cursing.Curse(name)
	fmt.Println(message1)
	fmt.Println(message2)
}
