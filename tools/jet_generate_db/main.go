package main

import (
	"github.com/go-jet/jet/v2/generator/postgres"
	_ "github.com/lib/pq"
)

func main() {

	pgConnection := postgres.DBConnection{
		Host:       "localhost",
		Port:       8032,
		User:       "postgres",
		Password:   "postgres",
		DBName:     "postgres",
		SchemaName: "public",
		SslMode:    "disable",
	}
	err := postgres.Generate("../../app/.jet_generated", pgConnection)
	if err != nil {
		panic(err)
	}
}
