package controller

import (
	middlewares "example.com/rest-api/Middlewares"
	"example.com/rest-api/service"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", service.GetEvents)
	server.GET("/events/:id", service.GetEvents)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", service.CreateEvent)
	authenticated.PUT("/events/:id", service.UpdateEvent)
	authenticated.DELETE("/events/:id", service.DeleteEvent)
	authenticated.POST("/events/:id/register", service.RegisterForEvent)
	authenticated.DELETE("/events/:id/register", service.CancelRegistration)

	server.POST("/signup", service.Signup)
	server.POST("/login", service.Login)
}
