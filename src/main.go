package main

import (
	"crud_vinos/pkg/infraestructure"
	"crud_vinos/pkg/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var db = *infraestructure.OpenDbConnection()

func main() {
	http.HandleFunc("/api/v1/vinos", handler)
	direccion := ":8000"
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Lógica para obtener vinos
		if id := r.URL.Query().Get("id"); id != "" {
			// Obtener un vino específico por ID
			vino, err := obtenerVinoPorID(&db, id)
			if err != nil {
				http.Error(w, "Vino no encontrado", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(vino)
		} else {
			// Obtener todos los vinos
			vinos, err := obtenerTodosLosVinos(&db)
			if err != nil {
				http.Error(w, "Error al obtener vinos", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(vinos)
		}
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func obtenerVinoPorID(db *sql.DB, id string) (*models.Vino, error) {
	query := "SELECT id, nombre, uva, pais FROM vinos WHERE id = ?"
	row := db.QueryRow(query, id)

	var vino models.Vino
	if err := row.Scan(&vino.ID, &vino.Nombre, &vino.Uva, &vino.Pais); err != nil {
		return nil, err
	}
	return &vino, nil
}

func obtenerTodosLosVinos(db *sql.DB) ([]models.Vino, error) {
	rows, err := db.Query("SELECT id, nombre, uva, pais FROM vinos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vinos []models.Vino
	for rows.Next() {
		var vino models.Vino
		if err := rows.Scan(&vino.ID, &vino.Nombre, &vino.Uva, &vino.Pais); err != nil {
			return nil, err
		}
		vinos = append(vinos, vino)
	}
	return vinos, nil
}
