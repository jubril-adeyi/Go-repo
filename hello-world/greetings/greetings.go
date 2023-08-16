package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name is given, return an error with message
	if name == "" {
		return "", errors.New("empty name")
	}
	// Return a greeting that embeds the name in a message.
	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
