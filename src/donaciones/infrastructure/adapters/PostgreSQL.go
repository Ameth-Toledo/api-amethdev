package adapters

import (
	"AmethToledo/src/donaciones/domain/entities"
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

// Save - Crear donación
func (pc *PostgreSQL) Save(donacion entities.Donacion) (entities.Donacion, error) {
	query := `INSERT INTO donaciones (usuario_id, modulo_id, monto, moneda, estado, metodo_pago, transaction_id, payment_id, fecha_pago) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`

	var id int
	err := pc.conn.QueryRow(
		query,
		donacion.UsuarioID,
		donacion.ModuloID,
		donacion.Monto,
		donacion.Moneda,
		donacion.Estado,
		donacion.MetodoPago,
		donacion.TransactionID,
		donacion.PaymentID,
		donacion.FechaPago,
	).Scan(&id)

	if err != nil {
		return entities.Donacion{}, fmt.Errorf("failed to save donacion: %v", err)
	}

	donacion.ID = id
	return donacion, nil
}

// GetById - Obtener donación por ID con nombres
func (pc *PostgreSQL) GetById(id int) (*entities.Donacion, error) {
	query := `SELECT 
              d.id, d.usuario_id, d.modulo_id, d.monto, d.moneda, d.estado, 
              d.metodo_pago, d.transaction_id, d.payment_id, d.fecha_pago,
              u.nombres, u.apellido_paterno, u.apellido_materno,
              m.titulo, c.nombre
              FROM donaciones d
              LEFT JOIN usuarios u ON d.usuario_id = u.id
              LEFT JOIN modulos m ON d.modulo_id = m.id
              LEFT JOIN cursos c ON m.id_curso = c.id
              WHERE d.id = $1`

	var donacion entities.Donacion
	var apellidoMaterno sql.NullString

	err := pc.conn.QueryRow(query, id).Scan(
		&donacion.ID,
		&donacion.UsuarioID,
		&donacion.ModuloID,
		&donacion.Monto,
		&donacion.Moneda,
		&donacion.Estado,
		&donacion.MetodoPago,
		&donacion.TransactionID,
		&donacion.PaymentID,
		&donacion.FechaPago,
		&donacion.UsuarioNombre,
		&donacion.UsuarioApellidos,
		&apellidoMaterno,
		&donacion.ModuloTitulo,
		&donacion.CursoNombre,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Donación no encontrada
		}
		return nil, fmt.Errorf("failed to get donacion by id: %v", err)
	}

	// Concatenar apellidos
	if apellidoMaterno.Valid {
		donacion.UsuarioApellidos += " " + apellidoMaterno.String
	}

	return &donacion, nil
}

// GetAll - Obtener todas las donaciones con nombres
func (pc *PostgreSQL) GetAll() ([]entities.Donacion, error) {
	query := `SELECT 
              d.id, d.usuario_id, d.modulo_id, d.monto, d.moneda, d.estado, 
              d.metodo_pago, d.transaction_id, d.payment_id, d.fecha_pago,
              u.nombres, 
              CONCAT(u.apellido_paterno, COALESCE(' ' || u.apellido_materno, '')) as apellidos,
              m.titulo, 
              c.nombre
              FROM donaciones d
              LEFT JOIN usuarios u ON d.usuario_id = u.id
              LEFT JOIN modulos m ON d.modulo_id = m.id
              LEFT JOIN cursos c ON m.id_curso = c.id
              ORDER BY d.fecha_pago DESC`

	rows, err := pc.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all donaciones: %v", err)
	}
	defer rows.Close()

	var donaciones []entities.Donacion
	for rows.Next() {
		var donacion entities.Donacion
		err := rows.Scan(
			&donacion.ID,
			&donacion.UsuarioID,
			&donacion.ModuloID,
			&donacion.Monto,
			&donacion.Moneda,
			&donacion.Estado,
			&donacion.MetodoPago,
			&donacion.TransactionID,
			&donacion.PaymentID,
			&donacion.FechaPago,
			&donacion.UsuarioNombre,
			&donacion.UsuarioApellidos,
			&donacion.ModuloTitulo,
			&donacion.CursoNombre,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan donacion: %v", err)
		}
		donaciones = append(donaciones, donacion)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return donaciones, nil
}

// GetByUsuarioID - Obtener donaciones por usuario con nombres
func (pc *PostgreSQL) GetByUsuarioID(usuarioID int) ([]entities.Donacion, error) {
	query := `SELECT 
              d.id, d.usuario_id, d.modulo_id, d.monto, d.moneda, d.estado, 
              d.metodo_pago, d.transaction_id, d.payment_id, d.fecha_pago,
              u.nombres, 
              CONCAT(u.apellido_paterno, COALESCE(' ' || u.apellido_materno, '')) as apellidos,
              m.titulo, 
              c.nombre
              FROM donaciones d
              LEFT JOIN usuarios u ON d.usuario_id = u.id
              LEFT JOIN modulos m ON d.modulo_id = m.id
              LEFT JOIN cursos c ON m.id_curso = c.id
              WHERE d.usuario_id = $1 
              ORDER BY d.fecha_pago DESC`

	rows, err := pc.conn.Query(query, usuarioID)
	if err != nil {
		return nil, fmt.Errorf("failed to get donaciones by usuario: %v", err)
	}
	defer rows.Close()

	var donaciones []entities.Donacion
	for rows.Next() {
		var donacion entities.Donacion
		err := rows.Scan(
			&donacion.ID,
			&donacion.UsuarioID,
			&donacion.ModuloID,
			&donacion.Monto,
			&donacion.Moneda,
			&donacion.Estado,
			&donacion.MetodoPago,
			&donacion.TransactionID,
			&donacion.PaymentID,
			&donacion.FechaPago,
			&donacion.UsuarioNombre,
			&donacion.UsuarioApellidos,
			&donacion.ModuloTitulo,
			&donacion.CursoNombre,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan donacion: %v", err)
		}
		donaciones = append(donaciones, donacion)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return donaciones, nil
}

// GetByModuloID - Obtener donaciones por módulo con nombres
func (pc *PostgreSQL) GetByModuloID(moduloID int) ([]entities.Donacion, error) {
	query := `SELECT 
              d.id, d.usuario_id, d.modulo_id, d.monto, d.moneda, d.estado, 
              d.metodo_pago, d.transaction_id, d.payment_id, d.fecha_pago,
              u.nombres, 
              CONCAT(u.apellido_paterno, COALESCE(' ' || u.apellido_materno, '')) as apellidos,
              m.titulo, 
              c.nombre
              FROM donaciones d
              LEFT JOIN usuarios u ON d.usuario_id = u.id
              LEFT JOIN modulos m ON d.modulo_id = m.id
              LEFT JOIN cursos c ON m.id_curso = c.id
              WHERE d.modulo_id = $1 
              ORDER BY d.fecha_pago DESC`

	rows, err := pc.conn.Query(query, moduloID)
	if err != nil {
		return nil, fmt.Errorf("failed to get donaciones by modulo: %v", err)
	}
	defer rows.Close()

	var donaciones []entities.Donacion
	for rows.Next() {
		var donacion entities.Donacion
		err := rows.Scan(
			&donacion.ID,
			&donacion.UsuarioID,
			&donacion.ModuloID,
			&donacion.Monto,
			&donacion.Moneda,
			&donacion.Estado,
			&donacion.MetodoPago,
			&donacion.TransactionID,
			&donacion.PaymentID,
			&donacion.FechaPago,
			&donacion.UsuarioNombre,
			&donacion.UsuarioApellidos,
			&donacion.ModuloTitulo,
			&donacion.CursoNombre,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan donacion: %v", err)
		}
		donaciones = append(donaciones, donacion)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return donaciones, nil
}

// GetByEstado - Obtener donaciones por estado con nombres
func (pc *PostgreSQL) GetByEstado(estado string) ([]entities.Donacion, error) {
	query := `SELECT 
              d.id, d.usuario_id, d.modulo_id, d.monto, d.moneda, d.estado, 
              d.metodo_pago, d.transaction_id, d.payment_id, d.fecha_pago,
              u.nombres, 
              CONCAT(u.apellido_paterno, COALESCE(' ' || u.apellido_materno, '')) as apellidos,
              m.titulo, 
              c.nombre
              FROM donaciones d
              LEFT JOIN usuarios u ON d.usuario_id = u.id
              LEFT JOIN modulos m ON d.modulo_id = m.id
              LEFT JOIN cursos c ON m.id_curso = c.id
              WHERE d.estado = $1 
              ORDER BY d.fecha_pago DESC`

	rows, err := pc.conn.Query(query, estado)
	if err != nil {
		return nil, fmt.Errorf("failed to get donaciones by estado: %v", err)
	}
	defer rows.Close()

	var donaciones []entities.Donacion
	for rows.Next() {
		var donacion entities.Donacion
		err := rows.Scan(
			&donacion.ID,
			&donacion.UsuarioID,
			&donacion.ModuloID,
			&donacion.Monto,
			&donacion.Moneda,
			&donacion.Estado,
			&donacion.MetodoPago,
			&donacion.TransactionID,
			&donacion.PaymentID,
			&donacion.FechaPago,
			&donacion.UsuarioNombre,
			&donacion.UsuarioApellidos,
			&donacion.ModuloTitulo,
			&donacion.CursoNombre,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan donacion: %v", err)
		}
		donaciones = append(donaciones, donacion)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return donaciones, nil
}

// GetTotal - Obtener total de donaciones
func (pc *PostgreSQL) GetTotal() (int, error) {
	query := `SELECT COUNT(*) FROM donaciones`

	var total int
	err := pc.conn.QueryRow(query).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total donaciones: %v", err)
	}

	return total, nil
}

// GetTotalByUsuario - Obtener total donado por usuario
func (pc *PostgreSQL) GetTotalByUsuario(usuarioID int) (float64, error) {
	query := `SELECT COALESCE(SUM(monto), 0) FROM donaciones WHERE usuario_id = $1 AND estado = 'completada'`

	var total float64
	err := pc.conn.QueryRow(query, usuarioID).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total by usuario: %v", err)
	}

	return total, nil
}

// GetTotalByModulo - Obtener total recaudado por módulo
func (pc *PostgreSQL) GetTotalByModulo(moduloID int) (float64, error) {
	query := `SELECT COALESCE(SUM(monto), 0) FROM donaciones WHERE modulo_id = $1 AND estado = 'completada'`

	var total float64
	err := pc.conn.QueryRow(query, moduloID).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("failed to get total by modulo: %v", err)
	}

	return total, nil
}

// Update
func (pc *PostgreSQL) Update(donacion entities.Donacion) error {
	query := `UPDATE donaciones SET 
              usuario_id = $1, 
              modulo_id = $2, 
              monto = $3, 
              moneda = $4, 
              estado = $5, 
              metodo_pago = $6, 
              transaction_id = $7, 
              payment_id = $8 
              WHERE id = $9`

	result, err := pc.conn.Exec(
		query,
		donacion.UsuarioID,     // $1
		donacion.ModuloID,      // $2
		donacion.Monto,         // $3
		donacion.Moneda,        // $4
		donacion.Estado,        // $5
		donacion.MetodoPago,    // $6
		donacion.TransactionID, // $7
		donacion.PaymentID,     // $8
		donacion.ID,            // $9
	)

	if err != nil {
		return fmt.Errorf("failed to update donacion: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("donacion not found")
	}

	return nil
}

// Delete - Eliminar donación
func (pc *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM donaciones WHERE id = $1`

	result, err := pc.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete donacion: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("donacion not found")
	}

	return nil
}
