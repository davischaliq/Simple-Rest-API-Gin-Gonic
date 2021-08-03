package config

import (
	"log"
	"os"

	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "123",
		Addr:     "localhost:5432",
		Database: "api",
	}
	var DB *pg.DB = pg.Connect(opts)

	if DB == nil {
		log.Printf("Failed to connect")
		os.Exit(100)
	} else {
		log.Printf("Connected to db")
	}
	return DB
}
