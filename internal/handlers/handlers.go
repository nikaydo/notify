package handlers

import (
	"main/internal/database"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Engine   *gin.Engine
	Database database.Database
}

func HandlersInit(e *gin.Engine, db database.Database) Handlers {
	Handler := Handlers{Engine: e, Database: db}
	api := Handler.Engine.Group("/api")

	users := Handler.Engine.Group("/user")
	users.POST("", Handler.createUser)

	events := api.Group("/event")
	events.POST("/create", Handler.createEvent)
	events.GET("/get", Handler.GetEvent)
	events.POST("/subscribe", Handler.SubscribeEvent)
	events.DELETE("/unsubscribe", Handler.UnSubscribeEvent)
	return Handler

}
