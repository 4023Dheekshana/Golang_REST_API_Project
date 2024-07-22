package routes

import (
	"dheek.com/restapi/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    //To get all the events
	server.GET("/events/:id", getEvent) //To get single event by id

	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)

	authenticated.POST("/events", createEvent)                       //To post the event in the database
	authenticated.PUT("/events/:id", updateEvent)                    //To update an event
	authenticated.DELETE("/events/:id", deleteEvent)                 //To delete an event
	authenticated.POST("/events/:id/register", resgisterForEvent)    // To register an event
	authenticated.DELETE("/events/:id/register", cancelRegistration) //To delete an registered event
	server.POST("/signup", signUp)                                   // To post user sign up details
	server.POST("/login", login)

}
