package handlers

import (
	"net/http"

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
	bookmarks, err := handler.repo.GetAll(ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, bookmarks)

}
