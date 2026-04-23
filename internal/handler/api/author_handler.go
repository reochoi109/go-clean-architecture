package api

import (
	"go-clean-architecture/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthorHandler struct {
	AuthorUsecase domain.AuthorUsecase
}

func NewAuthorHandler(group *gin.RouterGroup, uc domain.AuthorUsecase) {
	handler := &AuthorHandler{AuthorUsecase: uc}
	group.GET("/:id", handler.GetAuthorWithArticles)
}

func (h *AuthorHandler) GetAuthorWithArticles(c *gin.Context) {
	id := c.Param("id")
	author, articles, err := h.AuthorUsecase.GetAuthorWithArticles(c.Request.Context(), id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Author not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"author":   author,
		"articles": articles,
	})
}
