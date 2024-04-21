package handlers

import (
	"crud_vinos/internal/database"
	"crud_vinos/internal/repository"
	"encoding/json"
	"net/http"
)

var db = *database.OpenDbConnection()

func HandlerVino(w http.ResponseWriter, r *http.Request) {
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
				http.Error(w, "Error al buscar el vino", http.StatusInternalServerError)
				return
			} else if vino == nil {
				json.NewEncoder(w).Encode(map[string]interface{}{})
				return
			}
			json.NewEncoder(w).Encode(vino)
		} else {
			// Obtener todos los vinos
			vinos, err := repository.ObtenerTodosLosVinos(&db)
			if err != nil {
				http.Error(w, "Error al obtener vinos", http.StatusInternalServerError)
				return
			} else if vinos == nil {
				json.NewEncoder(w).Encode([]interface{}{})
				return
			}
			json.NewEncoder(w).Encode(vinos)
		}
	case "POST":
		// Lógica para crear un vino

		vino, err := repository.CrearVino(&db, r.Body)
		if err != nil {
			http.Error(w, "Error al crear vino", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(vino)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
