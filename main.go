package main

import (
	"flag"
	"go-linkr/internal/config"
	"go-linkr/internal/server"
	"log"
)

func main() {
	port := flag.String("port", "8383", "Port to listen on")
	dbPath := flag.String("db", "go-linkr-db.db", "Path to the sqlite3 database file")
	flag.Parse()

	cfg := config.NewConfig(*port, *dbPath)
	server := server.NewServer(cfg)

	log.Fatal(server.Run())
}
