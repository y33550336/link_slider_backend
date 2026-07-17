package main

import (
	"log"

	"github.com/y33550336/link_slider/config"
	"github.com/y33550336/link_slider/database"
	"github.com/y33550336/link_slider/handler"
	"github.com/y33550336/link_slider/repository"
)

func main() {
	cfg := &config.Config{}
	db, err := database.ConnectToDatabase(cfg.MySQLConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	server := handler.NewServer(cfg.AppAddr, repo)

	log.Println("Starting server on " + cfg.AppAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
