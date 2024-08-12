package greetings

import (
	"errors"
	"fmt"
)

/*
Hello returns and prints a greeting for the named person
Hello returns an empty string and and error if no name is given
*/
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
