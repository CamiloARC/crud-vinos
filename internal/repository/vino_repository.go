package repository

import (
	"crud_vinos/internal/models"
	"database/sql"
)

func ObtenerVinoPorID(db *sql.DB, id string) (*models.Vino, error) {
	query := "SELECT id, nombre, uva, pais FROM vinos WHERE id = ?"
	row := db.QueryRow(query, id)

	var vino models.Vino
	if err := row.Scan(&vino.ID, &vino.Nombre, &vino.Uva, &vino.Pais); err != nil {
		return nil, err
	}
	return &vino, nil
}

func ObtenerTodosLosVinos(db *sql.DB) ([]models.Vino, error) {
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
