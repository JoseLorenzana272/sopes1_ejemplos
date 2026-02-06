package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const TargetIP = "192.168.122.XX" // aqui cambian por la ip de su vm:)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://" + TargetIP + ":8082/datos")
		if err != nil {
			fmt.Fprintf(w, "Error contactando a Pong: %s", err)
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		fmt.Fprintf(w, "Ping desde Docker (VM1): La VM2 me respondió -> %s", string(body))
	})

	fmt.Println("API PING corriendo en puerto 8081")
	http.ListenAndServe(":8081", nil)
}
