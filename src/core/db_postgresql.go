package core

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Conn_PostgreSQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_PostgreSQL {
	if err := godotenv.Load(); err != nil {
		log.Printf("Advertencia: No se pudo cargar el archivo .env: %v", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbPort == "" {
		dbPort = "5432"
	}

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatalf("Error: Faltan variables de entorno. Verifica tu .env")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	db.SetMaxOpenConns(10)

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Error al verificar la conexión a la base de datos: %v", err)
	}

	fmt.Println("Conexión a PostgreSQL exitosa.")
	return &Conn_PostgreSQL{DB: db}
}

func (conn *Conn_PostgreSQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_PostgreSQL) FetchRows(query string, values ...interface{}) (*sql.Rows, error) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}

	return rows, nil
}
