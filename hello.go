package main

import (
	"fmt"
	"log"

	"github.com/jradhima/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := greetings.Hellos([]string{"yannis", "eleni"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
