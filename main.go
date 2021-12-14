package main

import (
	"API-Webservice/controllers"
	"API-Webservice/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/comments", controllers.FindComments)
	r.GET("/comments/:id", controllers.FindComment)
	r.POST("/comments", controllers.CreateComment)
	r.PATCH("/comments/:id", controllers.UpdateComment)
	r.DELETE("/comments/:id", controllers.DeleteComment)

	// Run the server
	r.Run()
}