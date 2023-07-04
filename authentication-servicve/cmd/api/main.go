package main

import (
	"authentication/data"
	"database/sql"
)

// webPort is the port for the web server.
const webPort = "80"

type Config struct {
	DB    *sql.DB
	Model data.model
}

func main() {}
