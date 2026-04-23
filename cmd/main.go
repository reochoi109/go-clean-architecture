package main

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/internal/infra/database"
	"go-clean-architecture/internal/repository/mysql"
	"go-clean-architecture/internal/usecase"

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

	authorRepo := mysql.NewAuthorRepository(db)
	articleRepo := mysql.NewArticleRepository(db)
	authorUsecase := usecase.NewAuthorUsecase(authorRepo, articleRepo)
}
