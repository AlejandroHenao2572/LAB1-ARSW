package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // Declara el WaitGroup globalmente

func countNumbers(start int, end int) {
	defer wg.Done() // Marca la goroutine como terminada al finalizar

	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func main() {
	wg.Add(3) // Indica que esperamos 3 goroutines

	go countNumbers(0, 99)
	go countNumbers(100, 199)
	go countNumbers(200, 299)

	wg.Wait() // Espera a que todas las goroutines terminen

	fmt.Println("Todas las goroutines terminaron")
}
