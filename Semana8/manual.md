
---

### **Objetivo Pedagógico**

1. **Señales:** Diferenciar entre un `SIGTERM` (petición amable) y un `SIGKILL` (asesinato).
2. **Concurrencia:** Ver múltiples "trabajadores" (Gorutinas) corriendo a la vez.
3. **Graceful Shutdown:** Aprender a limpiar recursos antes de morir (vital para el proyecto).

---

### **Paso a Paso (Manual para el Instructor)**

Haz esto en VS Code o en una terminal simple. No necesitas la VM si tienes Go instalado en tu host, pero si no, úsala.

#### **1. El Código (5 min)**

Crea un archivo llamado `inmortal.go`.
Este programa lanzará 3 trabajadores en paralelo y se negará a morir si presionas `Ctrl+C` una sola vez.

```go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Esta función simula un hilo o proceso trabajador
func trabajador(id int, stopCh chan bool) {
	for {
		select {
		case <-stopCh: // Si recibimos la orden de parar
			fmt.Printf("--> Trabajador %d: Deteniendo y guardando datos...\n", id)
			return
		default:
			// Simula trabajo
			fmt.Printf("[Trabajador %d] Procesando...\n", id)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	fmt.Printf("--- PID DEL PROCESO: %d ---\n", os.Getpid())
	fmt.Println("Intentar matarme con 'kill <PID>' o Ctrl+C")

	// 1. CANAL DE SEÑALES: Aquí interceptamos al Kernel
	// Creamos un canal para escuchar notificaciones del SO
	sigs := make(chan os.Signal, 1)
	
	// Le decimos al SO: "Si alguien manda SIGINT o SIGTERM, avísame a mí, no me mates"
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 2. CONCURRENCIA: Lanzamos gorutinas (Hilos ligeros)
	// Referencia a tu PDF: "Hilo vs Gorutina"
	stopCh := make(chan bool)
	go trabajador(1, stopCh)
	go trabajador(2, stopCh)
	go trabajador(3, stopCh)

	// 3. BLOQUEO: Esperamos la señal
	sig := <-sigs
	fmt.Printf("\n\n!!! ALERTA: Recibí la señal: %s !!!\n", sig)
	fmt.Println("Iniciando apagado ordenado (Graceful Shutdown)...")

	// Avisamos a los trabajadores que paren
	close(stopCh)
	
	// Simulamos tiempo de limpieza (cerrar BD, liberar memoria)
	time.Sleep(2 * time.Second)
	fmt.Println("Listo. He muerto en paz.")
}

```

#### **2. Ejecución y Prueba de Señales (5 min)**

Abre **dos terminales**.

**Terminal 1 (La Víctima):**
Ejecuta el programa:

```bash
go run inmortal.go

```

*Verás a los trabajadores imprimiendo mensajes.*

**Terminal 2 (El Verdugo):**
Aquí es donde enseñas los comandos de la Semana 8.

1. **Prueba 1: La Solicitud Amable (`SIGTERM`)**
* Busca el PID que imprimió el programa (ej. 5555).
* Ejecuta: `kill 5555` (Esto envía SIGTERM por defecto).
* **Resultado:** En la Terminal 1, el programa dirá: *"Recibí la señal... Iniciando apagado ordenado"* y terminará bien.
* **Lección:** "Esto es lo que pasa cuando Docker detiene un contenedor. Le da tiempo de guardar cosas."


2. **Prueba 2: La Interrupción (`SIGINT`)**
* Corre el programa de nuevo.
* Ve a la Terminal 1 y presiona `Ctrl + C`.
* **Resultado:** El programa lo detecta y se cierra limpiamente.
* **Lección:** "`Ctrl+C` no mata mágicamente, envía una señal `SIGINT` que el programa puede capturar."


3. **Prueba 3: El Asesinato (`SIGKILL`) - ¡La importante!**
* Corre el programa de nuevo.
* En la Terminal 2 ejecuta: `kill -9 5555` (o el PID actual).
* **Resultado:** En la Terminal 1 aparecerá "Killed" inmediatamente. **No** saldrá el mensaje de "Guardando datos".
* **Lección:** "La señal `-9` (SIGKILL) no se puede interceptar. El Kernel arranca el proceso de la CPU inmediatamente. Si estaban guardando en base de datos, esos datos se corrompieron."



---

#### **3. Hilos vs Gorutinas (Visualización con `htop`) (Opcional)**

Para cubrir el punto 5 de tu agenda ("Hilo vs Gorutina"):

1. Mientras corre el programa `inmortal.go`, ejecuta en la Terminal 2:
```bash
htop -p [PID]

```


*(Si no tienes htop: `top -H -p [PID]`)*.
2. **Observación:**
* Verás que aunque lanzamos 3 gorutinas, **no necesariamente ves 3 hilos del sistema operativo**.
* **Explicación:** "Aquí está la diferencia clave de la Semana 8. Un **Hilo POSIX (C)** es 1:1 con el Kernel (pesado). Una **Gorutina (Go)** es M:N (muchas gorutinas viven dentro de pocos hilos del SO). Por eso Go es ideal para servidores con miles de conexiones".

