package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HomePage Controller function displayes the homePage html template
func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", gin.H{
	})
}
