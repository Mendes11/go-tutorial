package main

import (
	"example.com/greetings"
	"fmt"
	"log"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message
	messages, err := greetings.Hellos([]string{"Mendes", "Mendoka", "Mendokão"})
	// If an error was returned, print it to the console and
	// exit the program
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console
	fmt.Println(messages)
	for key, value := range messages {
		fmt.Println(fmt.Sprintf("%s has a message of: %s", key, value))
	}
}
