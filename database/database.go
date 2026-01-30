package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	//open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	//Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	//set connettion pool setting
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database Connection succefully")
	return db, nil
}
