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

	//Request a greeting message
	message, err := greetings.Hello("")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

	// Get a quote
	fmt.Println(quote.Go())
}
