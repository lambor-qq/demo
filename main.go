package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	r.GET("/hello", func(c *gin.Context) {
		name := c.Query("name")
		if name == "" {
			name = "Lhb"
		}
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})
	})

	r.POST("/data", func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"received": data})
	})

	r.Run(":8081")
}