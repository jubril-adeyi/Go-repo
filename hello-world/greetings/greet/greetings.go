package greeting

import (
	// "errors"
	"errors"
	"fmt"
)

// function declaration (func) with the data
// type of the function parameter noted in there as string

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome !", name)
	return message, nil
}

func Greets(name string) string {

	message := fmt.Sprintf("Hi, %v. Welcome !", name)
	return message
}
