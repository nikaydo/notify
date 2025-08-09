package handlers

import (
	"main/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) createUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	uuid, err := h.Database.CreateUser(user)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(200, uuid.String())
}
