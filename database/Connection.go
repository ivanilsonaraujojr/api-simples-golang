package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	DB_HOST     = "localhost"
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_PORT     = 5432
	DB_NAME     = "apigo"
	MIGRATION   = `CREATE TABLE IF NOT EXISTS tarefas (
id serial PRIMARY KEY,
usuario text NOT NULL,
tarefa text NOT NULL,
dataCriacao timestamp with time zone DEFAULT current_timestamp,
concluida boolean DEFAULT false)
`
)

func OpenConnection() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(MIGRATION)
	if err != nil {
		log.Println("failed to run migratios")
		panic(err)
	}
	return db
}
