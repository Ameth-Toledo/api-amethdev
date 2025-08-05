package adapters

import (
	"AmethToledo/src/users/domain/entities"
	"database/sql"
	"fmt"
)

type PostgreSQL struct {
	conn *sql.DB
}

func NewPostgreSQL(conn *sql.DB) *PostgreSQL {
	return &PostgreSQL{conn: conn}
}

// Save - Crear usuario
func (pc *PostgreSQL) Save(user entities.User) (entities.User, error) {
	query := `INSERT INTO usuarios (nombres, apellido_paterno, apellido_materno, email, password_hash, rol_id, avatar, fecha_registro) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	var id int
	err := pc.conn.QueryRow(
		query,
		user.Nombres,
		user.ApellidoPaterno,
		user.ApellidoMaterno,
		user.Email,
		user.PasswordHash,
		user.RolID,
		user.Avatar,
		user.FechaRegistro,
	).Scan(&id)

	if err != nil {
		return entities.User{}, fmt.Errorf("failed to save user: %v", err)
	}

	user.ID = id
	return user, nil
}

// GetByCorreo - Buscar usuario por email
func (pc *PostgreSQL) GetByCorreo(email string) (*entities.User, error) {
	query := `SELECT id, nombres, apellido_paterno, apellido_materno, email, password_hash, rol_id, avatar, fecha_registro 
              FROM usuarios WHERE email = $1`

	var user entities.User
	err := pc.conn.QueryRow(query, email).Scan(
		&user.ID,
		&user.Nombres,
		&user.ApellidoPaterno,
		&user.ApellidoMaterno,
		&user.Email,
		&user.PasswordHash,
		&user.RolID,
		&user.Avatar,
		&user.FechaRegistro,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	return &user, nil
}

// GetAll - Obtener todos los usuarios
func (pc *PostgreSQL) GetAll() ([]entities.User, error) {
	query := `SELECT id, nombres, apellido_paterno, apellido_materno, email, password_hash, rol_id, avatar, fecha_registro 
              FROM usuarios ORDER BY fecha_registro DESC`

	rows, err := pc.conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %v", err)
	}
	defer rows.Close()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		err := rows.Scan(
			&user.ID,
			&user.Nombres,
			&user.ApellidoPaterno,
			&user.ApellidoMaterno,
			&user.Email,
			&user.PasswordHash,
			&user.RolID,
			&user.Avatar,
			&user.FechaRegistro,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	return users, nil
}

// GetById - Obtener usuario por ID
func (pc *PostgreSQL) GetById(id int) (*entities.User, error) {
	query := `SELECT id, nombres, apellido_paterno, apellido_materno, email, password_hash, rol_id, avatar, fecha_registro 
              FROM usuarios WHERE id = $1`

	var user entities.User
	err := pc.conn.QueryRow(query, id).Scan(
		&user.ID,
		&user.Nombres,
		&user.ApellidoPaterno,
		&user.ApellidoMaterno,
		&user.Email,
		&user.PasswordHash,
		&user.RolID,
		&user.Avatar,
		&user.FechaRegistro,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Usuario no encontrado
		}
		return nil, fmt.Errorf("failed to get user by id: %v", err)
	}

	return &user, nil
}

// Update - Actualizar usuario
func (pc *PostgreSQL) Update(user entities.User) error {
	query := `UPDATE usuarios SET 
              nombres = $2, 
              apellido_paterno = $3, 
              apellido_materno = $4, 
              email = $5, 
              rol_id = $6, 
              avatar = $7 
              WHERE id = $1`

	result, err := pc.conn.Exec(
		query,
		user.ID,
		user.Nombres,
		user.ApellidoPaterno,
		user.ApellidoMaterno,
		user.Email,
		user.RolID,
		user.Avatar,
	)

	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

// Delete - Eliminar usuario
func (pc *PostgreSQL) Delete(id int) error {
	query := `DELETE FROM usuarios WHERE id = $1`

	result, err := pc.conn.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
