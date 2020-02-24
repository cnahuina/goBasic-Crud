package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// getConnection obtiene una conexion a la BD

func getConnection() *sql.DB  {
	dsn := "postgres://postgres:123456@127.0.0.1:5432/dbApi?sslmode=disable"
	bd,err := sql.Open("postgres",dsn)
	if err != nil {
		log.Fatal(err)
	}

	return bd
}