# Laboratorio 1 ARSW - Concurrencia

**Arquitectura de Software**

**Autores:**
- David Alejandro Patacón Henao
- Daniel Felipe Hueso Rueda

---

## Descripción

Este laboratorio explora los conceptos fundamentales de **programación concurrente** mediante la creación de hilos que cuentan números en rangos específicos. El objetivo es entender la diferencia entre ejecución **concurrente** y **secuencial**, implementando la solución en dos lenguajes: **Java** y **Go**.

---

## Punto 1: Implementación de Hilos

### **Objetivo**
Crear 3 hilos que cuenten números en rangos diferentes:
- **Hilo 1:** 0 - 99
- **Hilo 2:** 100 - 199
- **Hilo 3:** 200 - 299

---

## Implementación en Java

### **Estructura del Código**

#### 1. Clase `CountThread.java`

**Características:**
- Extiende la clase `Thread` de Java
- Define rangos mediante los atributos `A` y `B`
- Sobrescribe el método `run()` para ejecutar el conteo

#### 2. Clase `CountThreadsMain.java`

### **Diferencia entre `start()` y `run()`**

#### **Método `start()`**
- **Concurrente:** Crea un nuevo hilo y ejecuta el método `run()` en paralelo
- **Comportamiento:** Los tres hilos se ejecutan simultáneamente
- **Resultado:** Salida intercalada y no ordenada
- **Uso:** Programación multihilo real  


#### **Método `run()`**
- **Secuencial:** Ejecuta el método directamente en el hilo actual (main)
- **Comportamiento:** Los hilos se ejecutan uno tras otro
- **Resultado:** Salida ordenada (0-299 en secuencia)
- **Uso:** No crea hilos nuevos, solo llama al método  

## Implementación en Go

#### 1. **`sync.WaitGroup`**
- Sincroniza la ejecución de goroutines
- `wg.Add(3)`: Indica que esperamos 3 goroutines
- `wg.Done()`: Marca una goroutine como finalizada
- `wg.Wait()`: Bloquea hasta que todas las goroutines terminen

#### 2. **Keyword `go`**
- Lanza una goroutine (equivalente a un hilo en Java)
- Ejecución concurrente automática
- Más eficiente que los hilos tradicionales

#### 3. **`defer wg.Done()`**
- Asegura que `Done()` se ejecute al finalizar la función  

## Análisis de Resultados

### **Ejecución Concurrente con `start()` en Java**
![Ejecución con start()](img/start.png)

**Observaciones:**
- Los números aparecen intercalados
- No hay orden predecible
- Los tres hilos compiten por la salida estándar
- Demuestra paralelismo real

### **Ejecución Secuencial con `run()` en Java**
![Ejecución con run()](img/run.png)

**Observaciones:**
- Los números aparecen en orden: 0-99, 100-199, 200-299
- Cada "hilo" termina antes de que comience el siguiente
- No hay concurrencia real, solo llamadas a métodos secuenciales
- Comportamiento similar a un bucle normal

### **Ejecución con Goroutines en Go**
![Ejecución con goroutines](img/ejecucionGo.png)

**Observaciones:**
- Salida intercalada similar a `start()` en Java
- Las goroutines se ejecutan concurrentemente
- `WaitGroup` asegura que el programa espere a todas las goroutines
- Mensaje final confirma la terminación de todas


---

## Punto 2: Ejercicio Black List Search

### Descripcion del Problema

El sistema debe validar si una dirección IP aparece en listas negras de servidores de manera concurrente. Se deben revisar 80,000 servidores distribuidos entre N hilos o goroutines, y reportar la IP como sospechosa si aparece en al menos 5 listas negras.

### Implementación en Java

#### Clase BlackListThread

Esta clase extiende `Thread` y se encarga de buscar una IP en un rango específico de servidores. Cada hilo trabaja de forma independiente en su rango asignado y mantiene un contador de servidores revisados y ocurrencias encontradas.

**Atributos principales:**
- `inicio` y `fin`: Definen el rango de servidores a revisar
- `ipaddress`: La IP que se está buscando
- `ocurrencesCount`: Número de veces que se encontró la IP
- `checkedServersCount`: Total de servidores revisados por este hilo

El método `run()` itera sobre el rango asignado, incrementa el contador de servidores revisados en cada iteración, y si encuentra la IP en un servidor, la agrega a la lista de ocurrencias.

#### Clase HostBlackListsValidator

Es el coordinador principal que distribuye el trabajo entre los hilos. Calcula cuántos servidores debe revisar cada hilo dividiendo el total entre N, y distribuye el resto(en los casos en donde N es impar)uniformemente entre los hilos para balancear la carga.

**Estrategia de distribución:**
Si hay 80,000 servidores y 4 hilos, cada hilo recibe exactamente 20,000 servidores.

Si hay 80,000 servidores y 3 hilos, cada hilo recibe aproximadamente 26,666 servidores(26,666*3 = 79,998) quedan sobrando 2 servidores. Los 2 servidores restantes se distribuyen dando uno extra a cada uno de los primeros 2 hilos.

Una vez creados los hilos, el método usa `join()` para esperar a que cada uno termine y luego recopila los resultados acumulando las ocurrencias y servidores revisados. Finalmente, si se encontraron 5 o más ocurrencias, reporta la IP como no confiable.

#### Clase Main

Crea una instancia del validador, especifica el número de hilos a utilizar y ejecuta la búsqueda para una IP específica.

### Implementación en Go

#### Archivo validator.go

Define la estructura `Validator` que coordina las goroutines y la estructura `SearchResult` que empaqueta los resultados de cada búsqueda.

El método `CheckHost` calcula la distribución de servidores por goroutine de forma similar a Java, lanzando N goroutines con la palabra clave `go`. Utiliza un `WaitGroup` para sincronizar la ejecución y un canal con buffer para recopilar los resultados.

Cada goroutine ejecuta la función `searchInRange`, que busca en su rango asignado y envía los resultados al canal. Al finalizar, el método principal recopila todos los resultados del canal y retorna la lista de servidores donde se encontró la IP.

#### Archivo datasource.go

**Nota:Este metodo fue generado por AI para simular la funcionalidad de la clase Java HostBlacklistsDataSourceFacade**
Implementa un Singleton usando `sync.Once` que simula la fuente de datos de listas negras, equivalente a la clase Java `HostBlacklistsDataSourceFacade`.

#### Archivo main.go

Crea el validador, especifica el número de goroutines y ejecuta la búsqueda para una IP de prueba.

### Comparación entre Java y Go

**Concurrencia:**
- Java usa hilos del sistema operativo creados con la clase `Thread` y el método `start()`
- Go usa goroutines que son más ligeras y se crean con la palabra clave `go`

**Sincronización:**
- Java espera a cada hilo individualmente usando `join()`
- Go usa `WaitGroup` para esperar a todas las goroutines como grupo

**Comunicación:**
- Java recopila resultados llamando getters después de `join()`
- Go usa canales para enviar resultados de forma concurrente

**Distribución de carga:**
Ambas implementaciones usan la misma estrategia de distribuir el resto uniformemente entre los hilos/goroutines para optimizar el balanceo de carga.

### Aspectos Técnicos Clave

**Balanceo de carga:** La distribucion uniforme del resto evita que un hilo termine mucho después que los demás, optimizando el tiempo total de ejecucion.

**Thread safety:** Cada hilo opera en un rango exclusivo de servidores, eliminando condiciones de carrera.


### Ejemplo de Ejecucion en Java
![Ejecución BlackListValidator Java](img/ejecucionJava.png)

### Ejemplo de Ejecucion en Go
![Ejecución BlackListValidator Go](img/ejecucionGoBL.png)




