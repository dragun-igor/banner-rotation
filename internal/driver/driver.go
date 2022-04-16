package driver

import (
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectPostgres(dsn string) (*DB, error) {
	d, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}
	dbConn.SQL = d
	err = testDB(d)

	return dbConn, err
}

func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error!", err)
	} else {
		log.Println("***Pinged database successfully!***")
	}
	return err
}
