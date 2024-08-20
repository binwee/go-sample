package routers

import (
	"github.com/binwee/go-sample/internal/handlers"
	"github.com/binwee/go-sample/internal/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	db *gorm.DB
	r  *gin.Engine
}

func NewRouter(db *gorm.DB, r *gin.Engine) Router {
	return Router{
		db: db,
		r:  r,
	}
}

func (router Router) RegisterRouter() {
	api := router.r.Group("/api")
	book := api.Group("/book")
	bookRepository := repositories.NewBookmarkRepository(router.db)
	bookHandler := handlers.NewBookHandler(bookRepository)
	book.GET("/getall", bookHandler.GetAll)
	book.GET("/getbyid", bookHandler.GetById)
	book.POST("/create", bookHandler.Create)
	book.PUT("/update", bookHandler.Update)
	book.DELETE("/delete", bookHandler.Delete)
}
