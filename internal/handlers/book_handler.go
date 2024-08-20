package handlers

import (
	"net/http"

	"github.com/binwee/go-sample/internal/models"
	"github.com/binwee/go-sample/internal/repositories"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	repo repositories.BookmarkRepository
}

func NewBookHandler(repo repositories.BookmarkRepository) *BookHandler {
	return &BookHandler{
		repo: repo,
	}
}

func (handler BookHandler) GetAll(c *gin.Context) {
	ctx := c.Request.Context()
	result, err := handler.repo.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (handler BookHandler) Create(c *gin.Context) {
	ctx := c.Request.Context()
	var book models.Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := handler.repo.Create(ctx, &book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
