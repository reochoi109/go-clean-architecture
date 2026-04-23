package main

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/internal/infra/database"
	"log"
)

func main() {
	cfg := config.Load()
	fmt.Printf("Loaded Config: %s\n", cfg)

	db, err := database.New(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
