package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"knative.dev/pkg/configmap"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/health", healthCheck)
	r.GET("/test", sampleApi)
	r.GET("/configmap", printConfigmap)
	r.Run("0.0.0.0:8080")
	log.SetOutput(os.Stdout)
}

func printConfigmap(c *gin.Context) {
	cfgMap, err := configmap.Load("config")
	if err != nil {
		fmt.Printf("error Occured while fetching properties from folder %v, with errror:%v", "config", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	props, _ := json.Marshal(ParseConfig(cfgMap["app.properties"]))
	c.JSON(http.StatusOK, string(props))
}

func sampleApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"test": "ok"})
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
