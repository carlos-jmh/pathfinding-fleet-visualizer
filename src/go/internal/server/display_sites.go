package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DisplayVisualizer displays the visualizer page.
func DisplayVisualizer(c *gin.Context) {
	c.HTML(http.StatusOK, "visualizer.tmpl", gin.H{
		"title": "Pathfinding Visualizer",
	})
}
