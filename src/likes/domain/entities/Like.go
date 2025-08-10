package entities

import "time"

type Like struct {
	ID              int       `json:"id" db:"id"`
	ModuloID        int       `json:"modulo_id" db:"modulo_id"`
	UsuarioID       *int      `json:"usuario_id,omitempty" db:"usuario_id"`
	FingerprintHash *string   `json:"fingerprint_hash,omitempty" db:"fingerprint_hash"`
	Fecha           time.Time `json:"fecha" db:"fecha"`
}
