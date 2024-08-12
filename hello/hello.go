package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"

	"log"
)

func main() {
	//Set the log entry prefix
	log.SetPrefix("greetings: ")
	//This flag disable printing tume, source file and line number
	log.SetFlags(0)

	// An array of names
	names := []string{"Felix", "Deborah", "Gladys"}

	//Request a greeting message
	messages, err := greetings.Hello(names[0])

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)

	// Get a quote
	fmt.Println(quote.Go())
}
