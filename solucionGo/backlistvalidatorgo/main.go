package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	testIP := "202.24.34.55"
	numCores := runtime.NumCPU()

	fmt.Printf("Nucleos: %d | IP: %s\n", numCores, testIP)

	// Configuraciones a probar (igual que en Java)
	numGoroutines := []int{1, numCores, numCores * 2, 50, 100, 1000, 100000}

	fmt.Println("\nGoroutines\tTiempo (ms)")
	fmt.Println("---------------------------")

	for _, n := range numGoroutines {
		validator := NewValidator()

		inicio := time.Now()
		validator.CheckHost(testIP, n)
		duracion := time.Since(inicio)

		fmt.Printf("%d\t\t%d\n", n, duracion.Milliseconds())
	}
}
