package main

import (
	"log"
	"net/http"
	"os"

	"letscode/project-02-slow-api/internal/api"
	"letscode/project-02-slow-api/internal/store"
)

func main() {
	path := os.Getenv("DB_PATH")
	if path == "" {
		path = "data/products.db"
	}
	db, err := store.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := store.Seed(db, 1000); err != nil {
		log.Fatal(err)
	}

	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("starting slow API on %s", addr)
	if err := http.ListenAndServe(addr, api.New(db)); err != nil {
		log.Fatal(err)
	}
}
