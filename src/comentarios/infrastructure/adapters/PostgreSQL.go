package adapters

import (
	"AmethToledo/src/comentarios/domain/entities"
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

// Save - Crear comentario
func (pc *PostgreSQL) Save(comentario entities.Comentario) (entities.Comentario, error) {
	query := `INSERT INTO comentarios (modulo_id, usuario_id, texto, fecha) 
              VALUES ($1, $2, $3, $4) RETURNING id`

	var id int
	err := pc.conn.QueryRow(
		query,
		comentario.ModuloID,
		comentario.UsuarioID,
		comentario.Texto,
		comentario.Fecha,
	).Scan(&id)

	if err != nil {
		return entities.Comentario{}, fmt.Errorf("failed to save comentario: %v", err)
	}

	comentario.ID = id
	return comentario, nil
}

// GetAll - Obtener todos los comentarios
func (pc *PostgreSQL) GetAll() ([]entities.Comentario, error) {
	query := `SELECT id, modulo_id, usuario_id, texto, fecha 
              FROM comentarios ORDER BY fecha DESC`

	rows, err := pc.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all comentarios: %v", err)
	}
	defer rows.Close()

	var comentarios []entities.Comentario
	for rows.Next() {
		var comentario entities.Comentario
		err := rows.Scan(
			&comentario.ID,
			&comentario.ModuloID,
			&comentario.UsuarioID,
			&comentario.Texto,
			&comentario.Fecha,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comentario: %v", err)
		}
		comentarios = append(comentarios, comentario)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return comentarios, nil
}

// GetById - Obtener comentario por ID
func (pc *PostgreSQL) GetById(id int) (*entities.Comentario, error) {
	query := `SELECT id, modulo_id, usuario_id, texto, fecha 
              FROM comentarios WHERE id = $1`

	var comentario entities.Comentario
	err := pc.conn.QueryRow(query, id).Scan(
		&comentario.ID,
		&comentario.ModuloID,
		&comentario.UsuarioID,
		&comentario.Texto,
		&comentario.Fecha,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Comentario no encontrado
		}
		return nil, fmt.Errorf("failed to get comentario by id: %v", err)
	}

	return &comentario, nil
}

// GetByModuloId - Obtener comentarios por ID de módulo
func (pc *PostgreSQL) GetByModuloId(moduloId int) ([]entities.Comentario, error) {
	query := `SELECT id, modulo_id, usuario_id, texto, fecha 
              FROM comentarios WHERE modulo_id = $1 ORDER BY fecha DESC`

	rows, err := pc.conn.Query(query, moduloId)
	if err != nil {
		return nil, fmt.Errorf("failed to get comentarios by modulo id: %v", err)
	}
	defer rows.Close()

	var comentarios []entities.Comentario
	for rows.Next() {
		var comentario entities.Comentario
		err := rows.Scan(
			&comentario.ID,
			&comentario.ModuloID,
			&comentario.UsuarioID,
			&comentario.Texto,
			&comentario.Fecha,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comentario: %v", err)
		}
		comentarios = append(comentarios, comentario)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return comentarios, nil
}

// GetByUsuarioId - Obtener comentarios por ID de usuario
func (pc *PostgreSQL) GetByUsuarioId(usuarioId int) ([]entities.Comentario, error) {
	query := `SELECT id, modulo_id, usuario_id, texto, fecha 
              FROM comentarios WHERE usuario_id = $1 ORDER BY fecha DESC`

	rows, err := pc.conn.Query(query, usuarioId)
	if err != nil {
		return nil, fmt.Errorf("failed to get comentarios by usuario id: %v", err)
	}
	defer rows.Close()

	var comentarios []entities.Comentario
	for rows.Next() {
		var comentario entities.Comentario
		err := rows.Scan(
			&comentario.ID,
			&comentario.ModuloID,
			&comentario.UsuarioID,
			&comentario.Texto,
			&comentario.Fecha,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comentario: %v", err)
		}
		comentarios = append(comentarios, comentario)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return comentarios, nil
}

// GetTotal - Obtener total de comentarios
func (pc *PostgreSQL) GetTotal() (int, error) {
	query := `SELECT COUNT(*) FROM comentarios`

	var total int
	err := pc.conn.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total comentarios: %v", err)
	}

	return total, nil
}

func (pc *PostgreSQL) Update(comentario entities.Comentario) error {
	query := `UPDATE comentarios SET 
              modulo_id = $1, 
              usuario_id = $2, 
              texto = $3
              WHERE id = $4`

	result, err := pc.conn.Exec(
		query,
		comentario.ModuloID,  // $1
		comentario.UsuarioID, // $2
		comentario.Texto,     // $3
		comentario.ID,        // $4
	)

	if err != nil {
		return fmt.Errorf("failed to update comentario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("comentario not found")
	}

	return nil
}

// Delete - Eliminar comentario
func (pc *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM comentarios WHERE id = $1`

	result, err := pc.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete comentario: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("comentario not found")
	}

	return nil
}

// GetByModuloIdWithUser - Obtener comentarios por ID de módulo con información del usuario
func (pc *PostgreSQL) GetByModuloIdWithUser(moduloId int) ([]entities.ComentarioConUsuario, error) {
	query := `SELECT c.id, c.modulo_id, c.usuario_id, c.texto, c.fecha,
                     u.nombres, u.apellido_paterno, u.apellido_materno, u.avatar
              FROM comentarios c
              JOIN usuarios u ON c.usuario_id = u.id
              WHERE c.modulo_id = $1 
              ORDER BY c.fecha DESC`

	rows, err := pc.conn.Query(query, moduloId)
	if err != nil {
		return nil, fmt.Errorf("failed to get comentarios by modulo id with user: %v", err)
	}
	defer rows.Close()

	var comentarios []entities.ComentarioConUsuario
	for rows.Next() {
		var comentario entities.ComentarioConUsuario
		err := rows.Scan(
			&comentario.ID,
			&comentario.ModuloID,
			&comentario.UsuarioID,
			&comentario.Texto,
			&comentario.Fecha,
			&comentario.NombreUsuario,
			&comentario.ApellidoPaterno,
			&comentario.ApellidoMaterno,
			&comentario.Avatar,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comentario with user: %v", err)
		}
		comentarios = append(comentarios, comentario)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return comentarios, nil
}

// GetAllWithUser - Obtener todos los comentarios con información del usuario
func (pc *PostgreSQL) GetAllWithUser() ([]entities.ComentarioConUsuario, error) {
	query := `SELECT c.id, c.modulo_id, c.usuario_id, c.texto, c.fecha,
                     u.nombres, u.apellido_paterno, u.apellido_materno, u.avatar
              FROM comentarios c
              JOIN usuarios u ON c.usuario_id = u.id
              ORDER BY c.fecha DESC`

	rows, err := pc.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all comentarios with user: %v", err)
	}
	defer rows.Close()

	var comentarios []entities.ComentarioConUsuario
	for rows.Next() {
		var comentario entities.ComentarioConUsuario
		err := rows.Scan(
			&comentario.ID,
			&comentario.ModuloID,
			&comentario.UsuarioID,
			&comentario.Texto,
			&comentario.Fecha,
			&comentario.NombreUsuario,
			&comentario.ApellidoPaterno,
			&comentario.ApellidoMaterno,
			&comentario.Avatar,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comentario with user: %v", err)
		}
		comentarios = append(comentarios, comentario)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return comentarios, nil
}
