package handlers

import (
	"crud_vinos/internal/database"
	"crud_vinos/internal/repository"
	"encoding/json"
	"net/http"
)

var db = *database.OpenDbConnection()

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Lógica para obtener vinos
		if id := r.URL.Query().Get("id"); id != "" {
			// Obtener un vino específico por ID
			vino, err := repository.ObtenerVinoPorID(&db, id)
			if err != nil {
				http.Error(w, "Vino no encontrado", http.StatusNotFound)
				return
			}
			json.NewEncoder(w).Encode(vino)
		} else {
			// Obtener todos los vinos
			vinos, err := repository.ObtenerTodosLosVinos(&db)
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
