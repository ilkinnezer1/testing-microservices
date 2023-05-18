package server

import (
	"authentication/data"
	"authentication/routes"
	"database/sql"
)

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func Start() {
	e := routes.AuthRoutes()

	if err := e.Start(":6000"); err != nil {
		panic(err)
	}
}
