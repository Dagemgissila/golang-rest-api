package routes

import (
	"github.com/gin-gonic/gin"
	"restapi.com/dagem/middleware"
)

func RegiterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authentication)
	authenticated.POST("events", createEvent)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
