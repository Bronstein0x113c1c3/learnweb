package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/test", func(ctx *gin.Context) {
		fmt.Print("sending")
		ctx.JSON(200, gin.H{
			"message": "OK!",
		})
	})
	server.Run("127.0.0.1:8080")
}
