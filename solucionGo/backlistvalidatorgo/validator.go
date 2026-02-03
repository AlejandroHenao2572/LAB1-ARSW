package main

import (
	"sync"
	"sync/atomic"
)

const BackListAlarmCount = 5

// Rsultado de la busqueda
type SearchResult struct {
	Occurrences     []int // Servidores donde se encontro la IP
	OcurrencesCount int   // Numero de ocurrencias encontradas
	checkedServes   int   // Numero de servidores revisados
}

// Valdator es el equivalente a HostBlackListsValidator en Java, valida las IPs en las listas negras
type Validator struct {
	dataSource             *BlacklistDataSource
	globalOccurrencesCount int32 // Contador global de ocurrencias, compartido entre goroutines
}

// NewValidator crea una nueva instancia de Validator
func NewValidator() *Validator {
	return &Validator{
		dataSource: GetInstance(),
	}
}

// CheckHost revisa si una IP esta en las listas negras usando n goroutines
// Retorna: lista de servidores donde se encontro la IP y total de servidores revisados
func (v *Validator) CheckHost(ip string, n int) ([]int, int) {
	totalServers := v.dataSource.GetRegisteredServersCount() // Total de servidores
	serversPerGoroutine := totalServers / n                  // Servidores por goroutine
	remainder := totalServers % n                            // Servidores restantes

	var wg sync.WaitGroup                    // WaitGroup para sincronizar goroutines
	wg.Add(n)                                // Esperar n goroutines
	resultChan := make(chan SearchResult, n) // Canal para resultados

	// Crear y lanzar goroutines
	for i := 0; i < n; i++ {
		// Asignar servidores a cada goroutine
		serversForThisGoroutine := serversPerGoroutine
		if i < remainder {
			serversForThisGoroutine++ // Distribuir los servidores restantes
		}
		start := i * serversPerGoroutine
		if i < remainder {
			start += i // Ajustar inicio para los servidores extra
		}
		end := start + serversForThisGoroutine

		//Lanzar la goroutine
		go v.searchInRange(start, end, ip, &wg, resultChan, &v.globalOccurrencesCount)
	}
	wg.Wait()         // Esperar a que todas las goroutines terminen
	close(resultChan) // Cerrar el canal de resultados

	//Recopilar resultados
	allOccurrences := []int{}
	totalOcurrencesCount := 0
	totalCheckedServers := 0

	for res := range resultChan {
		allOccurrences = append(allOccurrences, res.Occurrences...)
		totalOcurrencesCount += res.OcurrencesCount
		totalCheckedServers += res.checkedServes
	}

	return allOccurrences, totalCheckedServers
}

func (v *Validator) searchInRange(start, end int, ip string, wg *sync.WaitGroup, resultChan chan<- SearchResult, globalOccurrencesCount *int32) {
	defer wg.Done() // Marcar la goroutine como terminada al finalizar

	//Crear estructura de resultado
	result := SearchResult{
		Occurrences:     []int{},
		OcurrencesCount: 0,
		checkedServes:   0,
	}

	// Buscar en el rango asignado
	for i := start; i < end; i++ {

		if atomic.LoadInt32(globalOccurrencesCount) >= BackListAlarmCount {
			break // Salir del bucle si se ha alcanzado el umbral
		}

		result.checkedServes++
		// Verificar si la IP está en el servidor i
		if v.dataSource.IsInBlackListServer(i, ip) {
			result.Occurrences = append(result.Occurrences, i)
			result.OcurrencesCount++
			atomic.AddInt32(globalOccurrencesCount, 1)
		}
	}

	//Enviar resultado al canal
	resultChan <- result
}
