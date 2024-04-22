package handlers

import (
	"crud_vinos/internal/database"
	"crud_vinos/internal/models"
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
		if id := r.URL.Query().Get("id"); id != "" {
			vino, err := repository.ObtenerVinoPorID(&db, id)
			if err != nil {
				http.Error(w, "Error al buscar el vino", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(vino)
		} else {
			vinos, err := repository.ObtenerTodosLosVinos(&db)
			if err != nil {
				http.Error(w, "Error al obtener vinos", http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(vinos)
		}
	case "POST":
		var vino models.Vino
		if err := json.NewDecoder(r.Body).Decode(&vino); err != nil {
			http.Error(w, "Datos inválidos", http.StatusBadRequest)
			return
		}
		resp, err := repository.CrearVino(&db, &vino)
		if err != nil {
			http.Error(w, "Error al crear vino", http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(resp)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
