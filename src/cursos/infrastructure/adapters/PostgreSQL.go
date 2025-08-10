package adapters

import (
	"AmethToledo/src/cursos/domain/entities"
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

// Save - Crear curso
func (pc *PostgreSQL) Save(curso entities.Curso) (entities.Curso, error) {
	query := `INSERT INTO cursos (nombre, nivel, duracion, tecnologia, fecha, imagen, descripcion) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var id int
	err := pc.conn.QueryRow(
		query,
		curso.Nombre,
		curso.Nivel,
		curso.Duracion,
		curso.Tecnologia,
		curso.Fecha,
		curso.Imagen,
		curso.Descripcion,
	).Scan(&id)

	if err != nil {
		return entities.Curso{}, fmt.Errorf("failed to save curso: %v", err)
	}

	curso.ID = id
	return curso, nil
}

// GetAll - Obtener todos los cursos
func (pc *PostgreSQL) GetAll() ([]entities.Curso, error) {
	query := `SELECT id, nombre, nivel, duracion, tecnologia, fecha, imagen, descripcion 
              FROM cursos ORDER BY id DESC`

	rows, err := pc.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all cursos: %v", err)
	}
	defer rows.Close()

	var cursos []entities.Curso
	for rows.Next() {
		var curso entities.Curso
		err := rows.Scan(
			&curso.ID,
			&curso.Nombre,
			&curso.Nivel,
			&curso.Duracion,
			&curso.Tecnologia,
			&curso.Fecha,
			&curso.Imagen,
			&curso.Descripcion,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan curso: %v", err)
		}
		cursos = append(cursos, curso)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return cursos, nil
}

// GetById - Obtener curso por ID
func (pc *PostgreSQL) GetById(id int) (*entities.Curso, error) {
	query := `SELECT id, nombre, nivel, duracion, tecnologia, fecha, imagen, descripcion 
              FROM cursos WHERE id = $1`

	var curso entities.Curso
	err := pc.conn.QueryRow(query, id).Scan(
		&curso.ID,
		&curso.Nombre,
		&curso.Nivel,
		&curso.Duracion,
		&curso.Tecnologia,
		&curso.Fecha,
		&curso.Imagen,
		&curso.Descripcion,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Curso no encontrado
		}
		return nil, fmt.Errorf("failed to get curso by id: %v", err)
	}

	return &curso, nil
}

// GetByNombre - Buscar cursos por nombre
func (pc *PostgreSQL) GetByNombre(nombre string) ([]entities.Curso, error) {
	query := `SELECT id, nombre, nivel, duracion, tecnologia, fecha, imagen, descripcion 
              FROM cursos WHERE LOWER(nombre) LIKE LOWER($1) ORDER BY id DESC`

	rows, err := pc.conn.Query(query, "%"+nombre+"%")
	if err != nil {
		return nil, fmt.Errorf("failed to search cursos by name: %v", err)
	}
	defer rows.Close()

	var cursos []entities.Curso
	for rows.Next() {
		var curso entities.Curso
		err := rows.Scan(
			&curso.ID,
			&curso.Nombre,
			&curso.Nivel,
			&curso.Duracion,
			&curso.Tecnologia,
			&curso.Fecha,
			&curso.Imagen,
			&curso.Descripcion,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan curso: %v", err)
		}
		cursos = append(cursos, curso)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return cursos, nil
}

// GetTotal - Obtener total de cursos
func (pc *PostgreSQL) GetTotal() (int, error) {
	query := `SELECT COUNT(*) FROM cursos`

	var total int
	err := pc.conn.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total cursos: %v", err)
	}

	return total, nil
}

// Update - Actualizar curso
func (pc *PostgreSQL) Update(curso entities.Curso) error {
	query := `UPDATE cursos SET 
              nombre = $2, 
              nivel = $3, 
              duracion = $4, 
              tecnologia = $5, 
              fecha = $6, 
              imagen = $7, 
              descripcion = $8 
              WHERE id = $1`

	result, err := pc.conn.Exec(
		query,
		curso.ID,
		curso.Nombre,
		curso.Nivel,
		curso.Duracion,
		curso.Tecnologia,
		curso.Fecha,
		curso.Imagen,
		curso.Descripcion,
	)

	if err != nil {
		return fmt.Errorf("failed to update curso: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("curso not found")
	}

	return nil
}

// Delete - Eliminar curso
func (pc *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM cursos WHERE id = $1`

	result, err := pc.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete curso: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("curso not found")
	}

	return nil
}
