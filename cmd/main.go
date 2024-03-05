package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	g.Run(":3000")
}
