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

	occurrences, checkedServers := validator.CheckHost(testIP, n)

	fmt.Printf("\nThe host was found in the following blacklists: %v\n", occurrences)
	fmt.Printf("Checked %d servers out of %d\n", checkedServers, validator.dataSource.GetRegisteredServersCount())

	// Determinar si el host es confiable
	if len(occurrences) >= BackListAlarmCount {
		fmt.Printf("\nThe host %s is NOT trustworthy (found in %d blacklists)\n", testIP, len(occurrences))
	} else {
		fmt.Printf("\nThe host %s is trustworthy (found in %d blacklists)\n", testIP, len(occurrences))
	}
}
