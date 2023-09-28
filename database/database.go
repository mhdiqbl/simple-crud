package database

import (
	"database/sql"
	"log"
)

type DatabaseComponent struct {
	Master *sql.DB
}

func InitDatabase() *DatabaseComponent {
	var err error
	db := &DatabaseComponent{}

	db.Master, err = sql.Open("mysql", "root@tcp(localhost:3306)/otto_digital")
	db.Master.SetMaxIdleConns(10)
	db.Master.SetMaxOpenConns(20)

	if err != nil {
		log.Fatalf("Failed to bringing up mysql connection. err %v", err)
	}

	if err := db.Master.Ping(); err != nil {
		log.Fatalf("Failed to test mysql connection. err %v", err)
	}

	return db
}
