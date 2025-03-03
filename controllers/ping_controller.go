// controllers/ping_controller.go
package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// PingController handles ping requests
type PingController struct{}

// Ping returns a simple pong response
func (pc *PingController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
