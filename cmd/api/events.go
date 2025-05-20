package main

import (
	"event-api/internal/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (app *application) createEvent(c *gin.Context) {
	var event database.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := app.models.Events.Insert(&event)

	if err != nil {
		fmt.Println(err, "aaaa")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create event"})
	}

	c.JSON(http.StatusCreated, event)
}

func (app *application) getAllEvents(c *gin.Context) {
  events, err := app.models.Events.GetAll()

  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve events"})
  }

  c.JSON(http.StatusOK, events)
}

func (app *application) getEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})

	}

	event, err := app.models.Events.Get(id)

	if event == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
	}

	c.JSON(http.StatusOK, event)
}

func (app *application) updateEvent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
    return

	}

	existingEvent, err := app.models.Events.Get(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve event"})
    return
	}

	if existingEvent == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
	}

  updatedEvent := &database.Event{}

  if err := c.ShouldBindJSON(updatedEvent); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }

  updatedEvent.Id = id

	if err := app.models.Events.Update(updatedEvent); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update event"})
  }

  c.JSON(http.StatusOK, updatedEvent)
}

func (app *application) deleteEvent(c *gin.Context) {
  id, err := strconv.Atoi(c.Param("id"))

  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
  }

  if err := app.models.Events.Delete(id); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete event"})
  }

  c.JSON(http.StatusNoContent, nil)
}