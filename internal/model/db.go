package model

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Query struct {
	DB *sql.DB
}

func OpenDB() Query {
	db, err := sql.Open("postgres", "postgres://postgres:root@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Panicf("Failed To Connect to database %s \n", err)

	}
	createTable(db)

	return Query{DB: db}
}

type table struct {
	exist bool
}

func createTable(db *sql.DB) {
	stmt_create := `
	CREATE TABLE urls(
		ID  VARCHAR(7) PRIMARY KEY,
		url VARCHAR ,
		created_at timestamp default 'now()'
	)`
	var exists table
	log.Println("Checking if URLs table exists")
	result := db.QueryRow(`SELECT EXISTS (
		SELECT 1
		FROM pg_tables
		WHERE schemaname = 'public'
		AND tablename = 'urls'
	 );`)
	if err := result.Scan(&exists.exist); err != nil {
		log.Fatalf("Error has occurred %v", err)
	}
	if !exists.exist {
		log.Println("URLs table doesn't exist currently creating one")

		_, err := db.Exec(stmt_create)
		if err != nil {
			log.Fatalf("Error has occurred %v", err)

		} else {

			log.Println("Successfully Created Urls Table")

			return
		}
	}
	log.Println("Urls Table Exist")

}
