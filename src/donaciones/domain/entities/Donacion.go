package entities

import "time"

type Donacion struct {
	ID            int       `json:"id" db:"id"`
	UsuarioID     int       `json:"usuario_id" db:"usuario_id"`
	ModuloID      int       `json:"modulo_id" db:"modulo_id"`
	Monto         float64   `json:"monto" db:"monto"`
	Moneda        string    `json:"moneda" db:"moneda"`
	Estado        string    `json:"estado" db:"estado"`
	MetodoPago    string    `json:"metodo_pago" db:"metodo_pago"`
	TransactionID string    `json:"transaction_id" db:"transaction_id"`
	PaymentID     string    `json:"payment_id" db:"payment_id"`
	FechaPago     time.Time `json:"fecha_pago" db:"fecha_pago"`

	UsuarioNombre    string `json:"usuario_nombre,omitempty"`
	UsuarioApellidos string `json:"usuario_apellidos,omitempty"`
	ModuloTitulo     string `json:"modulo_titulo,omitempty"`
	CursoNombre      string `json:"curso_nombre,omitempty"`
}
