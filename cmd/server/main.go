package main

import (
	"database/sql"
	"github.com/MElghrbawy/threads/config"
	"log"

	"github.com/MElghrbawy/threads/internal/api"
	"github.com/MElghrbawy/threads/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	dbUrl := config.ConstructDatabaseURL(cfg)
	db, err := database.New(dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(db)

	app := fiber.New()
	api.SetupRoutes(app, db)

	serverAddr := config.ConstructServerAddress(cfg)
	log.Fatal(app.Listen(serverAddr))
}
