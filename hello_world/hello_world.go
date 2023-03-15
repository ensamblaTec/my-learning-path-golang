package main

import "fmt"

func main() {
	fmt.Println("Hello world from review code")
	fmt.Println(greeting("Misael"))
}

func greeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
