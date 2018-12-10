package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isayme/go-email/app/email"
	"github.com/isayme/go-email/app/manager"
)

// Send send email
func Send(c *gin.Context) {
	body := email.Message{}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	m := manager.Get()
	msgID, err := m.Sender.Send(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, msgID)
}
