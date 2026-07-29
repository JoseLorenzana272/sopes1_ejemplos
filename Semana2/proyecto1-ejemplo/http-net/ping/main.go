package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const TargetIP = "192.168.122.133"

func main() {
	http.HandleFunc("/iniciar", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://" + TargetIP + ":8082/responder")
		if err != nil {
			fmt.Fprintf(w, "Error contactando a Pong: %s", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Fprintf(w, "Ping: Llame a la otra API y me dijo: %s", string(body))
	})

	fmt.Println("API PING corriendo en puerto 8081")
	http.ListenAndServe(":8081", nil)
}
