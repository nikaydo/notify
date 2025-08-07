package handlers

import (
	"main/internal/database"
	"main/internal/models"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Engine   *gin.Engine
	Database database.Database
}

func HandlersInit(e *gin.Engine, db database.Database) Handlers {
	Handler := Handlers{Engine: e, Database: db}
	Handler.Engine.POST("/", Handler.createUser)
	return Handler

}

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
