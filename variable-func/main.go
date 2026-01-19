package main

import (
	"errors"
	"fmt"
	"log"
)

// Constants

const status string = "single"

func Somefunction(dating string) (string, error) {
	if dating == "" {
		return "relationship status :", errors.New(" but not single")
	}
	relationship_status := fmt.Sprintf("relationship status: %v ", dating)
	return relationship_status, nil
}

func main() {
	fmt.Println("main func here")
	// long way for variable and type decaration
	// var name string = "jbadeyi"
	// var age int = 20
	// var username = "jubril"
	// // username = "jbadeyi"
	// // name = "jubril"
	// // age = 20

	// // short way for variable and type decaration

	// name1 := "jayden"
	// age1 := 20
	// username1 := "jaydne"

	// fmt.Println(name, age, username, name1, age1, username1, status)

	log.SetPrefix("kinda complicated, ")
	log.SetFlags(0)

	real_status, err := Somefunction(status)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(real_status)
	}
}
