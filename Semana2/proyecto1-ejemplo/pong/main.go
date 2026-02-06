package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/datos", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Recibí petición de VM1")
		fmt.Fprintf(w, "Goku! Soy la API PONG respondiendo desde Containerd en VM2")
	})

	fmt.Println("API PONG corriendo en puerto 8082")
	http.ListenAndServe(":8082", nil)
}
