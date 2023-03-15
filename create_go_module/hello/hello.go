package main

import (
	"fmt"
	"log"

	"gcode/greetings"
)

func main() {
	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slices of names.
	names := []string{"Gladys's", "Samantha", "Darrin"}

	// Request a greeting messages for the names.
	// message, err := greetings.Hello("Pepe")
	message, err := greetings.Hellos(names)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If not error was returned, print returned message
	// to the console

	// Get a greeting message and print it.
	// message := greetings.Hello("Gladys")

	fmt.Println(message)
}

