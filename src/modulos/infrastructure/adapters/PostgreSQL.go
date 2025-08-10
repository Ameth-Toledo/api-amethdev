package adapters

import (
	"AmethToledo/src/modulos/domain/entities"
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

// Save - Crear módulo
func (pm *PostgreSQL) Save(modulo entities.Modulo) (entities.Modulo, error) {
	query := `INSERT INTO modulos (id_curso, imagen_portada, titulo, descripcion) 
              VALUES ($1, $2, $3, $4) RETURNING id`

	var id int
	err := pm.conn.QueryRow(
		query,
		modulo.IdCurso,
		modulo.ImagenPortada,
		modulo.Titulo,
		modulo.Descripcion,
	).Scan(&id)

	if err != nil {
		return entities.Modulo{}, fmt.Errorf("failed to save modulo: %v", err)
	}

	modulo.ID = id
	return modulo, nil
}

// GetAll - Obtener todos los módulos
func (pm *PostgreSQL) GetAll() ([]entities.Modulo, error) {
	query := `SELECT id, id_curso, imagen_portada, titulo, descripcion 
              FROM modulos ORDER BY id ASC`

	rows, err := pm.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all modulos: %v", err)
	}
	defer rows.Close()

	var modulos []entities.Modulo
	for rows.Next() {
		var modulo entities.Modulo
		err := rows.Scan(
			&modulo.ID,
			&modulo.IdCurso,
			&modulo.ImagenPortada,
			&modulo.Titulo,
			&modulo.Descripcion,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan modulo: %v", err)
		}
		modulos = append(modulos, modulo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return modulos, nil
}

// GetById - Obtener módulo por ID
func (pm *PostgreSQL) GetById(id int) (*entities.Modulo, error) {
	query := `SELECT id, id_curso, imagen_portada, titulo, descripcion 
              FROM modulos WHERE id = $1`

	var modulo entities.Modulo
	err := pm.conn.QueryRow(query, id).Scan(
		&modulo.ID,
		&modulo.IdCurso,
		&modulo.ImagenPortada,
		&modulo.Titulo,
		&modulo.Descripcion,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Módulo no encontrado
		}
		return nil, fmt.Errorf("failed to get modulo by id: %v", err)
	}

	return &modulo, nil
}

// GetByTitulo - Buscar módulos por título
func (pm *PostgreSQL) GetByTitulo(titulo string) ([]entities.Modulo, error) {
	query := `SELECT id, id_curso, imagen_portada, titulo, descripcion 
              FROM modulos WHERE LOWER(titulo) LIKE LOWER($1) ORDER BY id DESC`

	rows, err := pm.conn.Query(query, "%"+titulo+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to search modulos by titulo: %v", err)
	}
	defer rows.Close()

	var modulos []entities.Modulo
	for rows.Next() {
		var modulo entities.Modulo
		err := rows.Scan(
			&modulo.ID,
			&modulo.IdCurso,
			&modulo.ImagenPortada,
			&modulo.Titulo,
			&modulo.Descripcion,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan modulo: %v", err)
		}
		modulos = append(modulos, modulo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return modulos, nil
}

// Update - Actualizar módulo
func (pm *PostgreSQL) Update(modulo entities.Modulo) error {
	query := `UPDATE modulos SET 
              id_curso = $2, 
              imagen_portada = $3, 
              titulo = $4, 
              descripcion = $5 
              WHERE id = $1`

	result, err := pm.conn.Exec(
		query,
		modulo.ID,
		modulo.IdCurso,
		modulo.ImagenPortada,
		modulo.Titulo,
		modulo.Descripcion,
	)

	if err != nil {
		return fmt.Errorf("failed to update modulo: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("modulo not found")
	}

	return nil
}

// Delete - Eliminar módulo
func (pm *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM modulos WHERE id = $1`

	result, err := pm.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete modulo: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("modulo not found")
	}

	return nil
}
