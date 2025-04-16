package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ping request received")
		fmt.Fprintln(w, "Pong")
	})
	http.ListenAndServe(":8081", nil)
}
