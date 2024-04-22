package repository

import (
	"crud_vinos/internal/models"
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestObtenerVinoPorID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	id := "1"
	query := "SELECT id, nombre, uva, pais FROM vinos WHERE id = ?"

	// Configura la respuesta que esperas de la consulta
	rows := sqlmock.NewRows([]string{"id", "nombre", "uva", "pais"}).
		AddRow(id, "Frontera", "Cabernet", "Chile")

	// Espera que se ejecute la consulta con el ID proporcionado y responde con las filas definidas
	mock.ExpectQuery(query).
		WithArgs(id).
		WillReturnRows(rows)

	type args struct {
		db *sql.DB
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Vino
		wantErr bool
	}{
		{
			name: "Caso de prueba 1: Obtener vino con ID existente",
			args: args{
				db: db,
				id: "1",
			},
			want: &models.Vino{
				ID:     1,
				Nombre: "Frontera",
				Uva:    "Cabernet",
				Pais:   "Chile",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ObtenerVinoPorID(tt.args.db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ObtenerVinoPorID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ObtenerVinoPorID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCrearVino(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	vinoTest := &models.Vino{
		ID:     1,
		Nombre: "Frontera",
		Uva:    "Cabernet",
		Pais:   "Chile",
	}

	// Define la consulta de inserción que esperas ejecutar
	query := "INSERT INTO vinos (nombre, uva, pais) VALUES( ?, ?, ?)"

	// Simula un resultado de la consulta de inserción
	mock.ExpectExec(query).
		WithArgs(vinoTest.Nombre, vinoTest.Uva, vinoTest.Pais).
		WillReturnResult(sqlmock.NewResult(1, 1)) // id=1, 1 fila afectada

	type args struct {
		db   *sql.DB
		vino *models.Vino
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Vino
		wantErr bool
	}{
		{
			name: "Caso de prueba 2: Crear vino",
			args: args{
				db:   db,
				vino: vinoTest,
			},
			want:    vinoTest,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CrearVino(tt.args.db, tt.args.vino)
			if (err != nil) != tt.wantErr {
				t.Errorf("CrearVino() error = %v, wantErr ", err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CrearVino() = %v, want %v", got, tt.want)
			}
		})
	}
}
