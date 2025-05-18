package main

import (
 "net/http" 
 "github.com/gin-gonic/gin"
)

func (app *application) routes() http.Handler {
  g := gin.Default()

  v1 := g.Group("/api/v1") 
  {
    v1.POST("/events", app.createEvent)
    v1.GET("/events", app.getAllEvents)
    v1.GET("/events/:id", app.getEvent)
    v1.PUT("/events/:id", app.updateEvent)
    v1.DELETE("/events/:id", app.deleteEvent)
  }

  return g
}