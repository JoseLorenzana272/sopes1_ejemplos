package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// loggerMiddleware es un middleware que registra información sobre las peticiones HTTP
// next: El manejador HTTP que se ejecutará después del middleware
func loggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Registrar hora de inicio para calcular duración
		start := time.Now()

		// Loggear información de la petición entrante
		log.Printf("Inicio de petición %s %s", r.Method, r.URL.Path)

		// Llamar al siguiente manejador en la cadena
		next(w, r)

		// Loggear información de finalización con la duración
		log.Printf("Fin de petición %s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	}
}

// handler es la función que procesa las peticiones HTTP
// w: Writer para enviar la respuesta al cliente
// r: Request que contiene información sobre la petición HTTP
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola, este es un servidor Go!\n")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Todas las peticiones a "/" pasarán por loggerMiddleware primero
	http.HandleFunc("/", loggerMiddleware(handler))

	fmt.Println("Servidor escuchando en el puerto 80")

	// Iniciar el servidor HTTP en el puerto 80
	// ListenAndServe bloquea la ejecución hasta que ocurra un error
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
