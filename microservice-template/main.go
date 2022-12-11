package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/health", healthCheck)
	r.GET("/test", sampleApi)
	r.Run("0.0.0.0:8080")
}

func sampleApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "ok"})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
