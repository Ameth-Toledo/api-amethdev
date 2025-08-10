package domain

import (
	"AmethToledo/src/likes/domain/entities"
	"time"
)

type ModuloWithLikes struct {
	ModuloID  int `json:"modulo_id"`
	LikeCount int `json:"like_count"`
}

type LikeStats struct {
	ModuloID   int              `json:"modulo_id"`
	StartDate  time.Time        `json:"start_date"`
	EndDate    time.Time        `json:"end_date"`
	TotalLikes int              `json:"total_likes"`
	DailyStats []DailyLikeStat  `json:"daily_stats"`
	TopHours   []HourlyLikeStat `json:"top_hours"`
}

type DailyLikeStat struct {
	Date      time.Time `json:"date"`
	LikeCount int       `json:"like_count"`
}

type HourlyLikeStat struct {
	Hour      int `json:"hour"`
	LikeCount int `json:"like_count"`
}

type ILike interface {
	Save(like entities.Like) (entities.Like, error)
	GetByModulo(moduloID int) ([]entities.Like, error)
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
	GetLikesByFingerprint(fingerprint string) ([]entities.Like, error)
	GetLikeStatsByDateRange(moduloID int, startDate, endDate time.Time) (*LikeStats, error)
}
