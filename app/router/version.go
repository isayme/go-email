package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isayme/go-email/app"
)

// Version version info
func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    app.Name,
		"version": app.Version,
	})
}
