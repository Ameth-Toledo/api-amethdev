package adapters

import (
	"AmethToledo/src/contenidos/domain/entities"
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

// Save - Crear contenido
func (pc *PostgreSQL) Save(contenido entities.Contenido) (entities.Contenido, error) {
	query := `INSERT INTO contenidos (id_modulo, imagen_portada, titulo, descripcion, video_url, descripcion_module, repositorio) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	var id int
	err := pc.conn.QueryRow(
		query,
		contenido.IdModulo,
		contenido.ImagenPortada,
		contenido.Titulo,
		contenido.Descripcion,
		contenido.VideoURL,
		contenido.DescripcionModule,
		contenido.Repositorio,
	).Scan(&id)

	if err != nil {
		return entities.Contenido{}, fmt.Errorf("failed to save contenido: %v", err)
	}

	contenido.ID = id
	return contenido, nil
}

// GetAll - Obtener todos los contenidos
func (pc *PostgreSQL) GetAll() ([]entities.Contenido, error) {
	query := `SELECT id, id_modulo, imagen_portada, titulo, descripcion, video_url, descripcion_module, repositorio 
              FROM contenidos ORDER BY id ASC`

	rows, err := pc.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all contenidos: %v", err)
	}
	defer rows.Close()

	var contenidos []entities.Contenido
	for rows.Next() {
		var contenido entities.Contenido
		var repositorio sql.NullString

		err := rows.Scan(
			&contenido.ID,
			&contenido.IdModulo,
			&contenido.ImagenPortada,
			&contenido.Titulo,
			&contenido.Descripcion,
			&contenido.VideoURL,
			&contenido.DescripcionModule,
			&repositorio,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan contenido: %v", err)
		}

		// Manejar el campo nullable repositorio
		if repositorio.Valid {
			contenido.Repositorio = repositorio.String
		} else {
			contenido.Repositorio = ""
		}

		contenidos = append(contenidos, contenido)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return contenidos, nil
}

// GetById - Obtener contenido por ID
func (pc *PostgreSQL) GetById(id int) (*entities.Contenido, error) {
	query := `SELECT id, id_modulo, imagen_portada, titulo, descripcion, video_url, descripcion_module, repositorio 
              FROM contenidos WHERE id = $1`

	var contenido entities.Contenido
	var repositorio sql.NullString

	err := pc.conn.QueryRow(query, id).Scan(
		&contenido.ID,
		&contenido.IdModulo,
		&contenido.ImagenPortada,
		&contenido.Titulo,
		&contenido.Descripcion,
		&contenido.VideoURL,
		&contenido.DescripcionModule,
		&repositorio,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Contenido no encontrado
		}
		return nil, fmt.Errorf("failed to get contenido by id: %v", err)
	}

	// Manejar el campo nullable repositorio
	if repositorio.Valid {
		contenido.Repositorio = repositorio.String
	} else {
		contenido.Repositorio = ""
	}

	return &contenido, nil
}

// Update - Actualizar contenido
func (pc *PostgreSQL) Update(contenido entities.Contenido) error {
	query := `UPDATE contenidos SET 
              id_modulo = $2, 
              imagen_portada = $3, 
              titulo = $4, 
              descripcion = $5, 
              video_url = $6, 
              descripcion_module = $7, 
              repositorio = $8 
              WHERE id = $1`

	// Convertir repositorio vacío a NULL para la base de datos
	var repositorio interface{}
	if contenido.Repositorio == "" {
		repositorio = nil
	} else {
		repositorio = contenido.Repositorio
	}

	result, err := pc.conn.Exec(
		query,
		contenido.ID,
		contenido.IdModulo,
		contenido.ImagenPortada,
		contenido.Titulo,
		contenido.Descripcion,
		contenido.VideoURL,
		contenido.DescripcionModule,
		repositorio,
	)

	if err != nil {
		return fmt.Errorf("failed to update contenido: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("contenido not found")
	}

	return nil
}

// Delete - Eliminar contenido
func (pc *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM contenidos WHERE id = $1`

	result, err := pc.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete contenido: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("contenido not found")
	}

	return nil
}
