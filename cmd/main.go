package main

import (
	"fmt"
	"log"

	"github.com/binwee/go-sample/internal/config"
	"github.com/binwee/go-sample/internal/database"
	"github.com/binwee/go-sample/internal/handlers"
	"github.com/binwee/go-sample/internal/models"
	"github.com/binwee/go-sample/internal/repositories"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	db := database.GetDb(cfg)
	err = models.AutoMigrate(db)
	if err != nil {
		log.Fatal(err)
	}

	bookRepository := repositories.NewBookmarkRepository(db)
	bookHandler := handlers.NewBookHandler(bookRepository)

	r := gin.Default()
	r.GET("/api/getall", bookHandler.GetAll)
	r.Run(fmt.Sprintf(":%d", cfg.Server.Port))

}
