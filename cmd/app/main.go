package main

import (
	"crud_vinos/api/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/vinos", handlers.Handler)
	direccion := ":8000"
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}
