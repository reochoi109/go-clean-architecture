package main

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/internal/infra/database"
	"go-clean-architecture/internal/repository/mysql"
	"go-clean-architecture/internal/router"
	"go-clean-architecture/internal/usecase"

	"log"
)

func main() {
	// config
	cfg := config.Load()

	// database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Asia%%2FSeoul",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)
	db, err := database.New(cfg.Database.Type, dsn, database.Options{
		MaxOpenConns: cfg.Database.MaxOpenConns,
		MaxIdleConns: cfg.Database.MaxIdleConns,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// repository
	authorRepo := mysql.NewAuthorRepository(db)
	articleRepo := mysql.NewArticleRepository(db)

	// usecase
	ucs := &router.UsecaseContainer{
		AuthorUsecase: usecase.NewAuthorUsecase(authorRepo, articleRepo),
	}

	// router
	r := router.New(ucs)
	if err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatal(err)
	}
}
