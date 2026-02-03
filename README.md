# Laboratorio 1 ARSW - Concurrencia

**Arquitectura de Software**

**Autores:**
- David Alejandro Patacón Henao
- Daniel Felipe Hueso Rueda

---

## Descripción

Este laboratorio explora los conceptos fundamentales de **programación concurrente** mediante la creación de hilos que cuentan números en rangos específicos. El objetivo es entender la diferencia entre ejecución **concurrente** y **secuencial**, implementando la solución en dos lenguajes: **Java** y **Go**.

---

## Parte 1: Implementación de Hilos

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

## Parte 2: Ejercicio Black List Search

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

---

## Parte 2.1: Optimización - Detención Temprana

### Problema de Ineficiencia

La implementación actual presenta una limitacion: todos los hilos continuan buscando en sus rangos asignados incluso después de que colectivamente ya se han encontrado las 5 ocurrencias mínimas necesarias para reportar la IP como maliciosa. Esto significa que si las primeras ocurrencias se encuentran rápidamente, se siguen revisando innecesariamente miles de servidores adicionales.

**Ejemplo del problema:**
Si con 100 hilos las 5 ocurrencias se encuentran en los primeros 1,000 servidores, los 79,000 servidores restantes se siguen revisando de todas formas.

### Solución Propuesta: Detención Temprana

La optimización consiste en implementar un mecanismo que permita detener todos los hilos tan pronto como se alcance el umbral de 5 ocurrencias. Esto requiere:

**1. Variable compartida para el contador global de ocurrencias**

Los hilos deben poder consultar en tiempo real cuantas ocurrencias se han encontrado en total. Esto requiere una variable compartida accesible por todos los hilos.

**2. Sincronización del acceso a la variable compartida**

Como múltiples hilos leerán y modificarán el contador simultáneamente, se necesita sincronización para evitar condiciones de carrera. Las solciones para lenguaje son:

- **Java:** utilizar `AtomicInteger` para manejar el acceso a la variable compartida, en cada servidor encontrado se incrementa atomícamente. En cada iteración, el hilo verifica si el contador ha alcanzado el umbral necesario.

- **Go:** Usar `atomic` para proteger el acceso a la variable compartida. Cada goroutine incrementa el contador atómicamente y verifica su valor antes de continuar.

**3. Verificación periódica dentro del ciclo de búsqueda**

Cada hilo debe verificar regularmente si el contador global ha alcanzado 5 o la cantidad especificada de ocurrencias. Si es así, debe terminar su búsqueda inmediatamente.

**4. Notificación entre hilos**

Cuando un hilo encuentra una ocurrencia, debe actualizar el contador compartido y notificar a otros hilos que se alcanzó el límite.

### Nuevos Elementos que Introduce

**Variables compartidas:**
Ahora se requiere estado compartido entre hilos.

**Condiciones de carrera:**
Múltiples hilos accediendo y modificando la misma variable pueden causar inconsistencias si no se sincronizan correctamente.

**Overhead de sincronización:**
Verificar el contador compartido repetidamente y usar mecanismos de sincronización añade costo computacional. 

**Complejidad del código:**
La lógica se vuelve más compleja al manejar sincronización, verificaciones adicionales y terminación temprana de hilos.

**Ventajas:**
- Reduce el número de consultas cuando las ocurrencias se encuentran temprano
- Mejora el tiempo de respuesta en casos favorables
- Ahorra recursos de procesamiento

**Desventajas:**
- Mayor complejidad del código
- Posible contención si muchos hilos acceden al contador simultáneamente
- Requiere diseño cuidadoso para evitar errores de concurrencia

### Ejecución de la Optimización
En este caso se implementó la optimización tanto en Java como en Go, utilizando `AtomicInteger` en Java y `atomic` en Go para manejar el contador global de ocurrencias.

### Ejemplo de ejecucion Optimizada en Go
![alt text](img/ejecionGo2.png)

### Ejemplo de ejecucion Optimizada en Java
![alt text](img/ejecucionJava2.png)





