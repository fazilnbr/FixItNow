package db

import (
	"database/sql"
	"fmt"
	"log"

	config "github.com/fazilnbr/project-workey/pkg/config"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config) *sql.DB {

	databaseName := cfg.DBName

	dbURI := cfg.DBSOURCE

	//Opens database
	db, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}

	// verifies connection to the database is still alive
	err = db.Ping()
	if err != nil {
		fmt.Println("error in pinging")
		log.Fatal(err)

	}

	log.Println("\nConnected to database:", databaseName)

	return db

}
