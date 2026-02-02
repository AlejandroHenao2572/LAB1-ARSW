package main

import (
	"fmt"
	"sync"
)

//ESTA CLASE FUE GENERADA EN SU TOTALIDAD POR AI PARA SIMULAR LA FUNCIONALIDAD DE LA CLASE JAVA HostBlacklistsDataSourceFacade

// BlacklistDataSource simula la fuente de datos de listas negras
// Es equivalente a HostBlacklistsDataSourceFacade en Java
type BlacklistDataSource struct {
	// Map que contiene las ocurrencias: [serverID][IP] = true
	occurrences map[int]map[string]bool
	mutex       sync.RWMutex   // Para operaciones thread-safe de lectura/escritura
	threadHits  map[string]int // Contador de hits por goroutine (para debugging)
	hitsMutex   sync.Mutex     // Mutex para threadHits
}

var (
	instance *BlacklistDataSource
	once     sync.Once // Garantiza que solo se cree una instancia (Singleton)
)

// GetInstance retorna la única instancia del data source (patrón Singleton)
// Equivalente a getInstance() en Java
func GetInstance() *BlacklistDataSource {
	once.Do(func() {
		instance = &BlacklistDataSource{
			occurrences: make(map[int]map[string]bool),
			threadHits:  make(map[string]int),
		}
		instance.initializeData()
	})
	return instance
}

// initializeData inicializa los datos de prueba
// Replica exactamente los datos del bloque static en Java
func (ds *BlacklistDataSource) initializeData() {
	// Helper function para agregar ocurrencias fácilmente
	addOccurrence := func(server int, ip string) {
		if ds.occurrences[server] == nil {
			ds.occurrences[server] = make(map[string]bool)
		}
		ds.occurrences[server][ip] = true
	}

	// ========================================
	// IP encontrado rápido (200.24.34.55)
	// "to be found by a single thread"
	// ========================================
	addOccurrence(23, "200.24.34.55")
	addOccurrence(50, "200.24.34.55")
	addOccurrence(200, "200.24.34.55")
	addOccurrence(1000, "200.24.34.55")
	addOccurrence(500, "200.24.34.55")

	// ========================================
	// IP disperso (202.24.34.55)
	// "to be found through all threads"
	// ========================================
	addOccurrence(29, "202.24.34.55")
	addOccurrence(10034, "202.24.34.55")
	addOccurrence(20200, "202.24.34.55")
	addOccurrence(31000, "202.24.34.55")
	addOccurrence(70500, "202.24.34.55")

	// ========================================
	// IP para el tercer caso (202.24.34.54)
	// "to be found through all threads"
	// ========================================
	addOccurrence(39, "202.24.34.54")
	addOccurrence(10134, "202.24.34.54")
	addOccurrence(20300, "202.24.34.54")
	addOccurrence(70210, "202.24.34.54")

	fmt.Println("BlacklistDataSource initialized with test data")
}

// GetRegisteredServersCount retorna el total de servidores disponibles
// Equivalente a getRegisteredServersCount() en Java
func (ds *BlacklistDataSource) GetRegisteredServersCount() int {
	return 80000
}

// IsInBlackListServer verifica si una IP está en un servidor específico
// Equivalente a isInBlackListServer(int serverNumber, String ip) en Java
func (ds *BlacklistDataSource) IsInBlackListServer(serverNumber int, ip string) bool {
	// Trackear hits por goroutine (para debugging, como en Java)
	goroutineName := fmt.Sprintf("goroutine-%d", getGoroutineID())
	ds.trackThreadHit(goroutineName)

	// Lock de lectura (permite múltiples lecturas concurrentes)
	ds.mutex.RLock()
	defer ds.mutex.RUnlock()

	// Verificar si el servidor existe en el map
	serverMap, exists := ds.occurrences[serverNumber]
	if !exists {
		return false
	}

	// Verificar si la IP está en ese servidor
	return serverMap[ip]
}

// trackThreadHit registra un hit para una goroutine específica
func (ds *BlacklistDataSource) trackThreadHit(goroutineName string) {
	ds.hitsMutex.Lock()
	defer ds.hitsMutex.Unlock()

	ds.threadHits[goroutineName]++
}

// getGoroutineID obtiene el ID de la goroutine actual (simplificado)
// En producción, usarías runtime.Goexit() o similar
func getGoroutineID() int {
	// Simplificación: usar un contador atómico
	// En realidad, Go no expone IDs de goroutines fácilmente por diseño
	return 0 // Placeholder - puedes usar un contador si quieres
}

// ReportAsNotTrustworthy reporta una IP como no confiable
// Equivalente a reportAsNotTrustworthy(String host) en Java
func (ds *BlacklistDataSource) ReportAsNotTrustworthy(ip string) {
	fmt.Printf("HOST %s Reported as NOT trustworthy\n", ip)

	// Imprimir estadísticas de threads si está habilitado
	// (equivalente al System.getProperty("threadsinfo") en Java)
	ds.printThreadStats()
}

// ReportAsTrustworthy reporta una IP como confiable
// Equivalente a reportAsTrustworthy(String host) en Java
func (ds *BlacklistDataSource) ReportAsTrustworthy(ip string) {
	fmt.Printf("HOST %s Reported as trustworthy\n", ip)
}

// printThreadStats imprime estadísticas de uso de goroutines
func (ds *BlacklistDataSource) printThreadStats() {
	ds.hitsMutex.Lock()
	defer ds.hitsMutex.Unlock()

	if len(ds.threadHits) > 0 {
		fmt.Printf("\nTotal goroutines used: %d\n", len(ds.threadHits))
		fmt.Println("Hits per goroutine:")
		for name, hits := range ds.threadHits {
			fmt.Printf("  %s: %d hits\n", name, hits)
		}
		fmt.Println()
	}
}

// ResetStats limpia las estadísticas
func (ds *BlacklistDataSource) ResetStats() {
	ds.hitsMutex.Lock()
	defer ds.hitsMutex.Unlock()

	ds.threadHits = make(map[string]int)
}
