package models

type Vino struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Uva    string `json:"uva"`
	Pais   string `json:"pais"`
}
