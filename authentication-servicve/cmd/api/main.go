package main

import (
	"authentication/data"
	"database/sql"
)

const webPort = "80"

type Config struct {
	DB    *sql.DB
	Model data.model
}

func main() {}
