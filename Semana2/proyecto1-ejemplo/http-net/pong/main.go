package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/responder", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "¡Pong! (Desde Containerd en VM 2)")
	})

	fmt.Println("API PONG corriendo en puerto 8082")
	http.ListenAndServe(":8082", nil)
}
