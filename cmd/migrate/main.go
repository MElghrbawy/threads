package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MElghrbawy/print/config"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a Goose command")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx := context.Background()
	DBUrl := config.ConstructDatabaseURL(cfg)

	fmt.Println(DBUrl)
	db, err := sql.Open("postgres", DBUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(db)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Failed to set dialect: %v", err)
	}

	fmt.Println("migrate" + os.Args[1])
	migrationDir := os.Getenv("MIGRATION_DIR")
	fmt.Println(migrationDir)
	if err := goose.RunContext(ctx, os.Args[1], db, migrationDir, os.Args[2:]...); err != nil {
		log.Fatalf("Goose %s failed: %v", os.Args[1], err)
	}

}
