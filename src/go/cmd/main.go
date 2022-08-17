package main

import (
	"log"
	"net/http"

	"pathfinding-fleet-visualizer/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("../templates/**/*.tmpl")
	router.Static("/assets", "../assets")

	registerRoutes(router)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}

// registerRoutes registers the routes of the server
func registerRoutes(router *gin.Engine) {
	// always return 200 OK - for health checks
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.GET("/", server.DisplayVisualizer)
	router.POST("/dijkstra", server.GetPath)
}
