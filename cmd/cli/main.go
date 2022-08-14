package main

import (
	"database/sql"
	"github.com/bradleyshawkins/property/config"
	"github.com/bradleyshawkins/property/postgres"
	"github.com/bradleyshawkins/property/rest"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	conf, err := config.Load()
	if err != nil {
		log.Fatal("Unable to load config.", err)
	}

	db, err := sql.Open("postgres", conf.ConnectionString)
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	pg := postgres.NewDatabase(db)

	server := rest.NewServer(pg)

	log.Println(server)

}
