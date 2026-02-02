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