/// package main
package main

/// import fmt -> formatter
import "fmt"

/// import quote example
import "rsc.io/quote"

/// main function
func main() {
	fmt.Println("Hello world from review code")
	fmt.Println(greeting("Misael"))

	/// Call code in an external package
	fmt.Println(quote.Go())
}

/// greeting method with name parameter
func greeting(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

