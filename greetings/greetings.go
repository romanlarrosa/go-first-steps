package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

/*
Hello returns and prints a greeting for the named person
Hello returns an empty string and and error if no name is given
*/
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

/*
RandomFormat return one of a set of greeting messages.
The returned message is selected randomly
*/
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Returns a random format, usign rand module Intn function.
	return formats[rand.Intn(len(formats))]
}
