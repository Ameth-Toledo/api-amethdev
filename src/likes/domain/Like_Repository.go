package domain

import (
	"AmethToledo/src/likes/domain/entities"
	"time"
)

// Estructura extendida para mostrar información del módulo
type ModuloWithLikes struct {
	ModuloID     int    `json:"modulo_id"`
	ModuloTitulo string `json:"modulo_titulo"` // Nuevo campo
	LikeCount    int    `json:"like_count"`
}

// Estructura extendida para mostrar información del usuario en los likes
type LikeWithUserInfo struct {
	ID              int       `json:"id"`
	ModuloID        int       `json:"modulo_id"`
	ModuloTitulo    string    `json:"modulo_titulo"` // Nuevo campo
	UsuarioID       *int      `json:"usuario_id,omitempty"`
	UsuarioNombre   *string   `json:"usuario_nombre,omitempty"`   // Nuevo campo
	UsuarioApellido *string   `json:"usuario_apellido,omitempty"` // Nuevo campo
	FingerprintHash *string   `json:"fingerprint_hash,omitempty"`
	Fecha           time.Time `json:"fecha"`
}

// Estructura para el selector de módulos
type ModuloOption struct {
	ID     int    `json:"id"`
	Titulo string `json:"titulo"`
}

type LikeStats struct {
	ModuloID     int              `json:"modulo_id"`
	ModuloTitulo string           `json:"modulo_titulo"` // Nuevo campo
	StartDate    time.Time        `json:"start_date"`
	EndDate      time.Time        `json:"end_date"`
	TotalLikes   int              `json:"total_likes"`
	DailyStats   []DailyLikeStat  `json:"daily_stats"`
	TopHours     []HourlyLikeStat `json:"top_hours"`
}

type DailyLikeStat struct {
	Date      time.Time `json:"date"`
	LikeCount int       `json:"like_count"`
}

type HourlyLikeStat struct {
	Hour      int `json:"hour"`
	LikeCount int `json:"like_count"`
}

// Interfaz actualizada con nuevos métodos
type ILike interface {
	Save(like entities.Like) (entities.Like, error)
	GetByModulo(moduloID int) ([]entities.Like, error)
	GetByModuloWithUserInfo(moduloID int) ([]LikeWithUserInfo, error) // Nuevo método
	GetByModuloAndUser(moduloID int, usuarioID int) (*entities.Like, error)
	GetByModuloAndFingerprint(moduloID int, fingerprint string) (*entities.Like, error)
	Delete(id int) error
	DeleteByModuloAndUser(moduloID int, usuarioID int) error
	DeleteByModuloAndFingerprint(moduloID int, fingerprint string) error
	GetLikeCount(moduloID int) (int, error)
	CheckIfUserLiked(moduloID int, usuarioID int) (bool, error)
	CheckIfFingerprintLiked(moduloID int, fingerprint string) (bool, error)
	GetMostLikedModulos(limit int) ([]ModuloWithLikes, error)
	GetLikesByUser(usuarioID int) ([]entities.Like, error)
	GetLikesByUserWithModuleInfo(usuarioID int) ([]LikeWithUserInfo, error) // Nuevo método
	GetLikesByFingerprint(fingerprint string) ([]entities.Like, error)
	GetLikeStatsByDateRange(moduloID int, startDate, endDate time.Time) (*LikeStats, error)
	GetAllModulos() ([]ModuloOption, error) // Nuevo método para el selector
}
