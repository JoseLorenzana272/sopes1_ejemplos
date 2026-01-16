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
