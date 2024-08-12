package main

import (
	"fmt"

	"rsc.io/quote"

	"example.com/greetings"
)

func main() {
	fmt.Println(greetings.Hello("John"))
	fmt.Println(quote.Go())
}
