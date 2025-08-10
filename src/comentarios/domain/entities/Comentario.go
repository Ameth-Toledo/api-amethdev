package entities

import "time"

type Comentario struct {
	ID        int       `json:"id" db:"id"`
	ModuloID  int       `json:"modulo_id" db:"modulo_id"`
	UsuarioID int       `json:"usuario_id" db:"usuario_id"`
	Texto     string    `json:"texto" db:"texto"`
	Fecha     time.Time `json:"fecha" db:"fecha"`
}

type ComentarioConUsuario struct {
	ID              int       `json:"id" db:"id"`
	ModuloID        int       `json:"modulo_id" db:"modulo_id"`
	UsuarioID       int       `json:"usuario_id" db:"usuario_id"`
	Texto           string    `json:"texto" db:"texto"`
	Fecha           time.Time `json:"fecha" db:"fecha"`
	NombreUsuario   string    `json:"nombre_usuario" db:"nombre_usuario"`
	ApellidoPaterno string    `json:"apellido_paterno" db:"apellido_paterno"`
	ApellidoMaterno string    `json:"apellido_materno" db:"apellido_materno"`
	Avatar          int       `json:"avatar" db:"avatar"`
}
