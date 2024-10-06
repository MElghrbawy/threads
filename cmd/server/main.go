package main

import (
	"database/sql"
	"github.com/MElghrbawy/print/config"
	"log"

	"github.com/MElghrbawy/print/internal/api"
	"github.com/MElghrbawy/print/pkg/database"
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
