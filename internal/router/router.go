package router

import (
	"go-clean-architecture/domain"
	"go-clean-architecture/internal/handler/api"
	"go-clean-architecture/internal/router/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

const TEMP_TOKEN = "secret"

type UsecaseContainer struct {
	AuthorUsecase domain.AuthorUsecase
}

func New(uc *UsecaseContainer) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger())

	v1 := r.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"status":  "ok",
		})
	})

	authorGroup := v1.Group("/author")
	authorGroup.Use(middleware.AuthMiddleware(TEMP_TOKEN))
	{
		api.NewAuthorHandler(authorGroup, uc.AuthorUsecase)
	}
	return r
}
