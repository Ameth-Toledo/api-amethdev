package adapters

import (
	"AmethToledo/src/likes/domain"
	"AmethToledo/src/likes/domain/entities"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

// Save - Crear like
func (pl *PostgreSQL) Save(like entities.Like) (entities.Like, error) {
	query := `INSERT INTO likes (modulo_id, usuario_id, fingerprint_hash, fecha) 
              VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING id, fecha`

	err := pl.conn.QueryRow(
		query,
		like.ModuloID,
		like.UsuarioID,
		like.FingerprintHash,
	).Scan(&like.ID, &like.Fecha)

	if err != nil {
		return entities.Like{}, fmt.Errorf("failed to save like: %v", err)
	}

	return like, nil
}

// GetByModulo - Obtener todos los likes de un módulo
func (pl *PostgreSQL) GetByModulo(moduloID int) ([]entities.Like, error) {
	query := `SELECT id, modulo_id, usuario_id, fingerprint_hash, fecha 
              FROM likes WHERE modulo_id = $1 ORDER BY fecha DESC`

	rows, err := pl.conn.Query(query, moduloID)
	if err != nil {
		return nil, fmt.Errorf("failed to get likes by modulo: %v", err)
	}
	defer rows.Close()

	var likes []entities.Like
	for rows.Next() {
		var like entities.Like
		err := rows.Scan(
			&like.ID,
			&like.ModuloID,
			&like.UsuarioID,
			&like.FingerprintHash,
			&like.Fecha,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan like: %v", err)
		}
		likes = append(likes, like)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return likes, nil
}

// GetByModuloAndUser - Obtener like específico por módulo y usuario
func (pl *PostgreSQL) GetByModuloAndUser(moduloID int, usuarioID int) (*entities.Like, error) {
	query := `SELECT id, modulo_id, usuario_id, fingerprint_hash, fecha 
              FROM likes WHERE modulo_id = $1 AND usuario_id = $2`

	var like entities.Like
	err := pl.conn.QueryRow(query, moduloID, usuarioID).Scan(
		&like.ID,
		&like.ModuloID,
		&like.UsuarioID,
		&like.FingerprintHash,
		&like.Fecha,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Like no encontrado
		}
		return nil, fmt.Errorf("failed to get like by modulo and user: %v", err)
	}

	return &like, nil
}

// GetByModuloAndFingerprint - Obtener like específico por módulo y fingerprint
func (pl *PostgreSQL) GetByModuloAndFingerprint(moduloID int, fingerprint string) (*entities.Like, error) {
	query := `SELECT id, modulo_id, usuario_id, fingerprint_hash, fecha 
              FROM likes WHERE modulo_id = $1 AND fingerprint_hash = $2`

	var like entities.Like
	err := pl.conn.QueryRow(query, moduloID, fingerprint).Scan(
		&like.ID,
		&like.ModuloID,
		&like.UsuarioID,
		&like.FingerprintHash,
		&like.Fecha,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Like no encontrado
		}
		return nil, fmt.Errorf("failed to get like by modulo and fingerprint: %v", err)
	}

	return &like, nil
}

// Delete - Eliminar like por ID
func (pl *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM likes WHERE id = $1`

	result, err := pl.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete like: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("like not found")
	}

	return nil
}

// DeleteByModuloAndUser - Eliminar like por módulo y usuario
func (pl *PostgreSQL) DeleteByModuloAndUser(moduloID int, usuarioID int) error {
	query := `DELETE FROM likes WHERE modulo_id = $1 AND usuario_id = $2`

	result, err := pl.conn.Exec(query, moduloID, usuarioID)
	if err != nil {
		return fmt.Errorf("failed to delete like by modulo and user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("like not found")
	}

	return nil
}

// DeleteByModuloAndFingerprint - Eliminar like por módulo y fingerprint
func (pl *PostgreSQL) DeleteByModuloAndFingerprint(moduloID int, fingerprint string) error {
	query := `DELETE FROM likes WHERE modulo_id = $1 AND fingerprint_hash = $2`

	result, err := pl.conn.Exec(query, moduloID, fingerprint)
	if err != nil {
		return fmt.Errorf("failed to delete like by modulo and fingerprint: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("like not found")
	}

	return nil
}

// GetLikeCount - Obtener el número total de likes de un módulo
func (pl *PostgreSQL) GetLikeCount(moduloID int) (int, error) {
	query := `SELECT COUNT(*) FROM likes WHERE modulo_id = $1`

	var count int
	err := pl.conn.QueryRow(query, moduloID).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get like count: %v", err)
	}

	return count, nil
}

// CheckIfUserLiked - Verificar si un usuario ya dio like
func (pl *PostgreSQL) CheckIfUserLiked(moduloID int, usuarioID int) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM likes WHERE modulo_id = $1 AND usuario_id = $2)`

	var exists bool
	err := pl.conn.QueryRow(query, moduloID, usuarioID).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if user liked: %v", err)
	}

	return exists, nil
}

// CheckIfFingerprintLiked - Verificar si un fingerprint ya dio like
func (pl *PostgreSQL) CheckIfFingerprintLiked(moduloID int, fingerprint string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM likes WHERE modulo_id = $1 AND fingerprint_hash = $2)`

	var exists bool
	err := pl.conn.QueryRow(query, moduloID, fingerprint).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check if fingerprint liked: %v", err)
	}

	return exists, nil
}

// GetMostLikedModulos - Obtener módulos más likeados
func (pl *PostgreSQL) GetMostLikedModulos(limit int) ([]domain.ModuloWithLikes, error) {
	query := `
        SELECT modulo_id, COUNT(*) as like_count 
        FROM likes 
        GROUP BY modulo_id 
        ORDER BY like_count DESC 
        LIMIT $1`

	rows, err := pl.conn.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get most liked modulos: %v", err)
	}
	defer rows.Close()

	var modulos []domain.ModuloWithLikes
	for rows.Next() {
		var modulo domain.ModuloWithLikes
		err := rows.Scan(&modulo.ModuloID, &modulo.LikeCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan modulo with likes: %v", err)
		}
		modulos = append(modulos, modulo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return modulos, nil
}

// GetLikesByUser - Obtener todos los likes de un usuario
func (pl *PostgreSQL) GetLikesByUser(usuarioID int) ([]entities.Like, error) {
	query := `SELECT id, modulo_id, usuario_id, fingerprint_hash, fecha 
              FROM likes WHERE usuario_id = $1 ORDER BY fecha DESC`

	rows, err := pl.conn.Query(query, usuarioID)
	if err != nil {
		return nil, fmt.Errorf("failed to get likes by user: %v", err)
	}
	defer rows.Close()

	var likes []entities.Like
	for rows.Next() {
		var like entities.Like
		err := rows.Scan(
			&like.ID,
			&like.ModuloID,
			&like.UsuarioID,
			&like.FingerprintHash,
			&like.Fecha,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan like: %v", err)
		}
		likes = append(likes, like)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return likes, nil
}

// GetLikesByFingerprint - Obtener todos los likes de un fingerprint
func (pl *PostgreSQL) GetLikesByFingerprint(fingerprint string) ([]entities.Like, error) {
	query := `SELECT id, modulo_id, usuario_id, fingerprint_hash, fecha 
              FROM likes WHERE fingerprint_hash = $1 ORDER BY fecha DESC`

	rows, err := pl.conn.Query(query, fingerprint)
	if err != nil {
		return nil, fmt.Errorf("failed to get likes by fingerprint: %v", err)
	}
	defer rows.Close()

	var likes []entities.Like
	for rows.Next() {
		var like entities.Like
		err := rows.Scan(
			&like.ID,
			&like.ModuloID,
			&like.UsuarioID,
			&like.FingerprintHash,
			&like.Fecha,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan like: %v", err)
		}
		likes = append(likes, like)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return likes, nil
}

// GetLikeStatsByDateRange - Obtener estadísticas de likes por rango de fechas
func (pl *PostgreSQL) GetLikeStatsByDateRange(moduloID int, startDate, endDate time.Time) (*domain.LikeStats, error) {
	// Obtener total de likes en el período
	totalQuery := `SELECT COUNT(*) FROM likes 
                   WHERE modulo_id = $1 AND fecha >= $2 AND fecha <= $3`

	var totalLikes int
	err := pl.conn.QueryRow(totalQuery, moduloID, startDate, endDate).Scan(&totalLikes)
	if err != nil {
		return nil, fmt.Errorf("failed to get total likes: %v", err)
	}

	// Obtener estadísticas diarias
	dailyQuery := `
        SELECT DATE(fecha)::text as date, COUNT(*) as count 
        FROM likes 
        WHERE modulo_id = $1 AND fecha >= $2 AND fecha <= $3 
        GROUP BY DATE(fecha) 
        ORDER BY DATE(fecha)`

	dailyRows, err := pl.conn.Query(dailyQuery, moduloID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily stats: %v", err)
	}
	defer dailyRows.Close()

	var dailyStats []domain.DailyLikeStat
	for dailyRows.Next() {
		var stat domain.DailyLikeStat
		var dateStr string
		err := dailyRows.Scan(&dateStr, &stat.LikeCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan daily stat: %v", err)
		}

		// Parsear la fecha con manejo flexible
		stat.Date, err = parseFlexibleDBDate(dateStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %v", err)
		}

		dailyStats = append(dailyStats, stat)
	}

	// Obtener estadísticas por hora (top 10 horas)
	hourlyQuery := `
        SELECT EXTRACT(HOUR FROM fecha) as hour, COUNT(*) as count 
        FROM likes 
        WHERE modulo_id = $1 AND fecha >= $2 AND fecha <= $3 
        GROUP BY EXTRACT(HOUR FROM fecha) 
        ORDER BY count DESC 
        LIMIT 10`

	hourlyRows, err := pl.conn.Query(hourlyQuery, moduloID, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get hourly stats: %v", err)
	}
	defer hourlyRows.Close()

	var hourlyStats []domain.HourlyLikeStat
	for hourlyRows.Next() {
		var stat domain.HourlyLikeStat
		err := hourlyRows.Scan(&stat.Hour, &stat.LikeCount)
		if err != nil {
			return nil, fmt.Errorf("failed to scan hourly stat: %v", err)
		}
		hourlyStats = append(hourlyStats, stat)
	}

	return &domain.LikeStats{
		ModuloID:   moduloID,
		StartDate:  startDate,
		EndDate:    endDate,
		TotalLikes: totalLikes,
		DailyStats: dailyStats,
		TopHours:   hourlyStats,
	}, nil
}

// parseFlexibleDBDate maneja diferentes formatos de fecha que pueden venir de la DB
func parseFlexibleDBDate(dateStr string) (time.Time, error) {
	// Limpiar espacios en blanco
	dateStr = strings.TrimSpace(dateStr)

	if dateStr == "" {
		return time.Time{}, fmt.Errorf("empty date string")
	}

	// Lista de formatos que PostgreSQL podría devolver
	formats := []string{
		"2006-01-02",                 // YYYY-MM-DD (esperado)
		"2006-01-02T15:04:05Z",       // ISO con Z
		"2006-01-02T15:04:05-07:00",  // ISO con timezone
		"2006-01-02T15:04:05",        // ISO sin timezone
		"2006-01-02 15:04:05",        // YYYY-MM-DD HH:MM:SS
		"2006-01-02 15:04:05.999999", // Con microsegundos
		"2006-01-02 15:04:05-07",     // Con timezone
		time.RFC3339,                 // RFC3339 completo
	}

	var lastErr error
	for _, format := range formats {
		parsed, err := time.Parse(format, dateStr)
		if err == nil {
			// Para fechas diarias, normalizar a medianoche UTC
			if format == "2006-01-02" {
				return time.Date(parsed.Year(), parsed.Month(), parsed.Day(), 0, 0, 0, 0, time.UTC), nil
			}
			return parsed, nil
		} else {
			lastErr = err
		}
	}

	return time.Time{}, fmt.Errorf("unsupported date format '%s': %v", dateStr, lastErr)
}
