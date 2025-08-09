package handlers

import (
	"main/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handlers) createEvent(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	uuid, err := h.Database.CreateEvent(event)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.String(201, uuid.String())
}

func (h *Handlers) GetEvent(c *gin.Context) {
	var event models.ReturnValues
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	e, err := h.Database.GetEvent(event.UuidEvent)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			c.String(400, "event with this uuid not found")
			return
		}
		c.String(500, err.Error())
		return
	}
	c.JSON(200, e)
}

func (h *Handlers) SubscribeEvent(c *gin.Context) {
	var event models.ReturnValues
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.Database.SubscribeEvent(event.UuidUser, event.UuidEvent)
	if err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			c.String(409, "user alredy subscribed")
			return
		}
		c.String(500, err.Error())
		return
	}
	c.Status(201)
}

func (h *Handlers) UnSubscribeEvent(c *gin.Context) {
	var event models.ReturnValues
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.Database.UnSubscribeEvent(event.UuidUser, event.UuidEvent)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.Status(200)
}
