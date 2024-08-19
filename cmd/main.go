package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/binwee/go-sample/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.GetConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.GET("/hello", hello)
	log.Fatal(r.Run(fmt.Sprintf(":%d", cfg.Server.Port)))
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
