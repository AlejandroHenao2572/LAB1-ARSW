package main

import (
	"fmt"
)

func main() {
	validator := NewValidator()

	// NUmero de goroutines a utilizar
	n := 100

	// Probar con el IP del ejemplo de java
	testIP := "200.24.34.55"

	occurrences := validator.CheckHost(testIP, n)

	fmt.Printf("\nThe host was found in the following blacklists: %v\n", occurrences)
}
