package main

import (
	"crud_vinos/api/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/vinos", handlers.HandlerVino)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API de vinos")
	})
	direccion := ":8080"
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
