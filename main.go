package main

import (
	"log"

	"syncvault/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	_, err := config.ConectarBanco()
	if err != nil {
		log.Fatal("não conectou no banco: ", err)
	}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
