package entities

import "time"

type User struct {
	ID              int       `json:"id" db:"id"`
	Nombres         string    `json:"nombres" db:"nombres"`
	ApellidoPaterno string    `json:"apellido_paterno" db:"apellido_paterno"`
	ApellidoMaterno string    `json:"apellido_materno" db:"apellido_materno"`
	Email           string    `json:"email" db:"email"`
	PasswordHash    string    `json:"-" db:"password_hash"`
	RolID           int       `json:"rol_id" db:"rol_id"`
	Avatar          int       `json:"avatar" db:"avatar"`
	FechaRegistro   time.Time `json:"fecha_registro" db:"fecha_registro"`
}
