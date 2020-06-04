package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// StatusController struct
type StatusController struct{}

// Status ...
func (s StatusController) Status(c *gin.Context) {
	c.String(http.StatusOK, "pong!")
}
