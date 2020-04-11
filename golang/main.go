package main

import (
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong", "go_version": runtime.GOOS})
	})
	r.Run(":5000")
}
