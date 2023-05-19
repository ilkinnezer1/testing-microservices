package server

import (
	"authentication/data"
	"authentication/routes"
	"database/sql"
	"log"
	"os"
	"time"
)

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)

	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Database is not ready")
			counts++
		} else {
			log.Println("Connected to Database")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for 2 seconds")
		time.Sleep(time.Second * 2)
		continue
	}
}

func Start() {
	e := routes.AuthRoutes()

	conn := connectToDB()
	if conn == nil {
		log.Panicln("Can't connect to the postgres")
	}

	_ = Config{
		DB:     conn,
		Models: data.New(conn),
	}

	if err := e.Start(":6000"); err != nil {
		panic(err)
	}
}
